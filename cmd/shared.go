package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/jippi/scm-engine/pkg/config"
	"github.com/jippi/scm-engine/pkg/integration/backstage"
	"github.com/jippi/scm-engine/pkg/scm"
	"github.com/jippi/scm-engine/pkg/scm/github"
	"github.com/jippi/scm-engine/pkg/scm/gitlab"
	"github.com/jippi/scm-engine/pkg/state"
	"github.com/teris-io/shortid"
	slogctx "github.com/veqryn/slog-context"
)

var sid = shortid.MustNew(1, shortid.DefaultABC, 2342)

func getClient(ctx context.Context) (scm.Client, error) {
	backstageClient, err := backstage.NewClient(ctx, state.BackstageURL(ctx), state.BackstageToken(ctx), nil)
	if err != nil {
		slogctx.Warn(ctx, "Backstage client is not available, actions requiring it will be skipped", slog.Any("error", err))
	}

	switch state.Provider(ctx) {
	case "github":
		return github.NewClient(ctx), nil

	case "gitlab":
		return gitlab.NewClient(ctx, backstageClient)

	default:
		return nil, fmt.Errorf("unknown provider %q - we only support 'github' and 'gitlab'", state.Provider(ctx))
	}
}

func ProcessMR(ctx context.Context, client scm.Client, cfg *config.Config, event any) (err error) {
	// Track start time of the evaluation
	ctx = state.WithStartTime(ctx, time.Now())

	// Attach unique eval id to the logs so they are easy to filter on later
	ctx = state.WithEvaluationID(ctx, sid.MustGenerate())

	// Track where we grab the configuration file from
	ctx = slogctx.With(ctx, slog.String("config_source_branch", "merge_request_branch"))

	// Should we allow failing the CI pipeline?
	allowPipelineFailure := false

	defer state.LockForProcessing(ctx)()

	// Stop the pipeline when we leave this func
	defer func() {
		if stopErr := client.Stop(ctx, err, allowPipelineFailure); stopErr != nil {
			slogctx.Error(ctx, "Failed to update pipeline", slog.Any("error", stopErr))
		}
	}()

	// Start the pipeline
	if err := client.Start(ctx); err != nil {
		return fmt.Errorf("failed to update pipeline monitor: %w", err)
	}

	//
	// Create and validate the evaluation context
	//

	slogctx.Info(ctx, "Creating evaluation context")

	evalContext, err := client.EvalContext(ctx)
	if err != nil {
		return err
	}

	if evalContext == nil || !evalContext.IsValid() {
		slogctx.Warn(ctx, "Evaluating context is empty, does the Merge Request exists?")

		return nil
	}

	// Check if we are allowed to fail the CI pipeline
	allowPipelineFailure = evalContext.AllowPipelineFailure(ctx)

	//
	// (Optional) Download the .scm-engine.yml configuration file from the GitLab HTTP API
	//

	var (
		configShouldBeDownloaded = cfg == nil
		configSourceRef          = state.CommitSHA(ctx)
	)

	// If the current branch is not in a state where the config file can be trusted,
	// we instead use the HEAD version of the file
	if !evalContext.CanUseConfigurationFileFromChangeRequest(ctx) {
		configShouldBeDownloaded = true
		configSourceRef = "HEAD"

		// Update the logger with new value
		ctx = slogctx.With(ctx, slog.String("config_source_branch", configSourceRef))
	}

	// Download and parse the configuration file if necessary
	if configShouldBeDownloaded {
		slogctx.Debug(ctx, "Downloading scm-engine configuration from ref: "+configSourceRef)

		file, err := client.MergeRequests().GetRemoteConfig(ctx, state.ConfigFilePath(ctx), configSourceRef)
		if err != nil {
			slogctx.Warn(ctx, "Could not read remote config file", slog.Any("error", err))
		} else {
			// Parse the file
			cfg, err = config.ParseFile(file)
			if err != nil { // error on parsing failures when present
				return fmt.Errorf("could not parse config file: %w", err)
			}
		}
	}

	// Merge previously loaded config with Repository config
	globalConfig := config.GlobalConfigFromContext(ctx) // the global config if previously loaded
	if globalConfig != nil && cfg != nil {
		cfg = globalConfig.Merge(cfg)
	}

	// Sanity check for having a configuration loaded
	if cfg == nil {
		return errors.New("cfg==nil; this is unexpected an error, please report!")
	}

	// Load any remote configuration files
	if err := cfg.LoadIncludes(ctx, client); err != nil {
		return fmt.Errorf("failed to load 'include' settings: %w", err)
	}

	// Allow changing the 'dry-run' mode via configuration file
	if cfg.DryRun != nil && *cfg.DryRun != state.IsDryRun(ctx) {
		slogctx.Info(ctx, "Configuration file has a 'dry_run' value, using that in favor of server default")

		ctx = state.WithDryRun(ctx, *cfg.DryRun)
	}

	// Lint the configuration file to catch any misconfigurations
	if err := cfg.Lint(ctx, evalContext); err != nil {
		return fmt.Errorf("Configuration failed validation: %w", err)
	}

	// Write the config to context so we can pull it out later
	// If a global config file was set, this overrides the global config with the merged global and repository config
	ctx = config.WithConfig(ctx, cfg)

	//
	// Do the actual context evaluation
	//

	slogctx.Info(ctx, "Evaluating context")

	evalContext.SetWebhookEvent(event)
	evalContext.SetContext(ctx)

	labels, actions, err := cfg.Evaluate(ctx, evalContext)
	if err != nil {
		return err
	}

	slogctx.Debug(ctx, "Evaluation complete", slog.Int("number_of_labels", len(labels)), slog.Int("number_of_actions", len(actions)))

	//
	// Post-evaluation sync of labels
	//

	slogctx.Info(ctx, "Sync labels")

	if err := syncLabels(ctx, client, labels); err != nil {
		return err
	}

	var (
		add    scm.LabelOptions
		remove scm.LabelOptions
	)

	existingLabels := evalContext.GetLabels()

	for _, e := range labels {
		if e.Matched && !slices.Contains(existingLabels, e.Name) {
			add = append(add, e.Name)
		} else if !e.Matched && slices.Contains(existingLabels, e.Name) {
			remove = append(remove, e.Name)
		}
	}

	//
	// Post-evaluation sync of actions
	//

	update := &scm.UpdateMergeRequestOptions{}

	if len(add) > 0 {
		update.AddLabels = &add
	}

	if len(remove) > 0 {
		update.RemoveLabels = &remove
	}

	slogctx.Info(ctx, "Applying actions")

	if err := runActions(ctx, evalContext, client, update, actions); err != nil {
		return err
	}

	//
	// Update the Merge Request with the outcome of labels and actions
	//

	slogctx.Info(ctx, "Updating Merge Request")

	return updateMergeRequest(ctx, client, update)
}

func updateMergeRequest(ctx context.Context, client scm.Client, update *scm.UpdateMergeRequestOptions) error {
	if update == nil || reflect.DeepEqual(update, &scm.UpdateMergeRequestOptions{}) {
		slogctx.Info(ctx, "No changes to apply to Merge Request")

		return nil
	}

	if state.IsDryRun(ctx) {
		slogctx.Info(ctx, "In dry-run, dumping the update struct we would send to GitLab", slog.Any("changes", update))

		return nil
	}

	slogctx.Debug(ctx, "Applying Merge Request changes", slog.Any("changes", update))

	_, err := client.MergeRequests().Update(ctx, update)

	return err
}

func runActions(ctx context.Context, evalContext scm.EvalContext, client scm.Client, update *scm.UpdateMergeRequestOptions, actions config.Actions) error {
	if len(actions) == 0 {
		slogctx.Debug(ctx, "No actions evaluated to true, skipping")

		return nil
	}

	for _, action := range actions {
		ctx := slogctx.With(ctx, slog.String("action_name", action.Name))
		slogctx.Info(ctx, "Applying action")

		if evalContext.HasExecutedActionGroup(action.Group) {
			slogctx.Warn(ctx, fmt.Sprintf("Already executed another action within group '%s'; skipping current action until next evaluation", action.Group))

			continue
		}

		evalContext.TrackActionGroupExecution(action.Group)

		for _, task := range action.Then {
			if err := client.ApplyStep(ctx, evalContext, update, task); err != nil {
				slogctx.Error(ctx, "failed to apply action step", slog.Any("error", err))

				return err
			}
		}
	}

	return nil
}

func syncLabels(ctx context.Context, client scm.Client, required []scm.EvaluationResult) error {
	slogctx.Info(ctx, "Going to sync required labels", slog.Int("number_of_labels", len(required)))

	remote, err := client.Labels().List(ctx)
	if err != nil {
		return err
	}

	remoteLabels := map[string]*scm.Label{}
	for _, e := range remote {
		remoteLabels[e.Name] = e
	}

	// Create
	for _, label := range required {
		if _, ok := remoteLabels[label.Name]; ok {
			continue
		}

		slogctx.Info(ctx, "Creating label", slog.String("label", label.Name))

		if state.IsDryRun(ctx) {
			continue
		}

		_, resp, err := client.Labels().Create(ctx, &scm.CreateLabelOptions{
			Name:        &label.Name,        //nolint:gosec
			Color:       &label.Color,       //nolint:gosec
			Description: &label.Description, //nolint:gosec
			Priority:    label.Priority,
		})
		if err != nil {
			// Label already exists
			if resp.StatusCode == http.StatusConflict {
				slogctx.Warn(ctx, "Label already exists", slog.String("label", label.Name))

				continue
			}

			return err
		}
	}

	// Update
	for _, label := range required {
		remote, ok := remoteLabels[label.Name]
		if !ok {
			continue
		}

		if label.IsEqual(ctx, remote) {
			continue
		}

		slogctx.Info(ctx, "Updating label", slog.String("label", label.Name))

		if state.IsDryRun(ctx) {
			continue
		}

		_, _, err := client.Labels().Update(ctx, &scm.UpdateLabelOptions{
			Name:        &label.Name,        //nolint:gosec
			Color:       &label.Color,       //nolint:gosec
			Description: &label.Description, //nolint:gosec
			Priority:    label.Priority,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
