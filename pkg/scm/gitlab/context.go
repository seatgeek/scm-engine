package gitlab

import (
	"context"
	"log/slog"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/jippi/scm-engine/pkg/scm"
	"github.com/jippi/scm-engine/pkg/state"
	slogctx "github.com/veqryn/slog-context"
	"golang.org/x/oauth2"
)

var _ scm.EvalContext = (*Context)(nil)

// pipelineByIDResponse is used for a separate low-complexity query to fetch pipeline by ID (pipeline events only).
// It requests only scalar pipeline fields (no jobs) to stay under GitLab's complexity limit.
type pipelineByIDResponse struct {
	Project *struct {
		Pipeline *minimalPipelineForQuery `graphql:"pipeline(iid: $pipeline_id)"`
	} `graphql:"project(fullPath: $project_id)"`
}

type minimalPipelineForQuery struct {
	Active        bool               `graphql:"active"`
	Cancelable    bool               `graphql:"cancelable"`
	Complete      bool               `graphql:"complete"`
	Duration      *int               `graphql:"duration"`
	FailureReason *string            `graphql:"failureReason"`
	FinishedAt    *time.Time         `graphql:"finishedAt"`
	ID            string             `graphql:"id"`
	Iid           string             `graphql:"iid"`
	Latest        bool               `graphql:"latest"`
	Name          *string            `graphql:"name"`
	Path          *string            `graphql:"path"`
	Retryable     bool               `graphql:"retryable"`
	StartedAt     *time.Time         `graphql:"startedAt"`
	Status        PipelineStatusEnum `graphql:"status"`
	Stuck         bool               `graphql:"stuck"`
	TotalJobs     int                `graphql:"totalJobs"`
	UpdatedAt     time.Time          `graphql:"updatedAt"`
	Warnings      bool               `graphql:"warnings"`
}

func (m *minimalPipelineForQuery) toContextPipeline() *ContextPipeline {
	if m == nil {
		return nil
	}
	return &ContextPipeline{
		Active:        m.Active,
		Cancelable:    m.Cancelable,
		Complete:      m.Complete,
		Duration:      m.Duration,
		FailureReason: m.FailureReason,
		FinishedAt:    m.FinishedAt,
		ID:            m.ID,
		Iid:           m.Iid,
		Latest:        m.Latest,
		Name:          m.Name,
		Path:          m.Path,
		Retryable:     m.Retryable,
		StartedAt:     m.StartedAt,
		Status:        m.Status,
		Stuck:         m.Stuck,
		TotalJobs:     m.TotalJobs,
		UpdatedAt:     m.UpdatedAt,
		Warnings:      m.Warnings,
		Jobs:         nil, // not fetched in minimal query to keep complexity low
	}
}

// fetchPipelineByID runs a small GraphQL query to load a pipeline by ID. Used only for pipeline webhook events.
func fetchPipelineByID(ctx context.Context, client *graphql.Client, projectID, pipelineID string) *ContextPipeline {
	var resp pipelineByIDResponse
	err := client.Query(ctx, &resp, map[string]any{
		"project_id":  graphql.ID(projectID),
		"pipeline_id": graphql.ID(pipelineID),
	})
	if err != nil {
		slogctx.Debug(ctx, "Failed to fetch pipeline by ID (pipeline event)", slog.Any("error", err))
		return nil
	}
	if resp.Project == nil || resp.Project.Pipeline == nil {
		return nil
	}
	return resp.Project.Pipeline.toContextPipeline()
}

func NewContext(ctx context.Context, baseURL, token string) (*Context, error) {
	httpClient := oauth2.NewClient(
		ctx,
		oauth2.StaticTokenSource(
			&oauth2.Token{
				AccessToken: token,
			},
		),
	)

	client := graphql.NewClient(baseURL+"/api/graphql", httpClient)

	var (
		evalContext *Context
		variables   = map[string]any{
			"project_id": graphql.ID(state.ProjectID(ctx)),
			"mr_id":      state.MergeRequestID(ctx),
		}
	)

	if err := client.Query(ctx, &evalContext, variables); err != nil {
		return nil, err
	}

	if evalContext.Project == nil || evalContext.Project.MergeRequest == nil {
		return nil, nil //nolint:nilnil
	}

	// Initialize null-able types
	evalContext.ActionGroups = make(map[string]any)

	// Move project labels into a un-nested expr exposed field
	evalContext.Project.Labels = evalContext.Project.ResponseLabels.Nodes
	evalContext.Project.ResponseLabels.Nodes = nil

	// Move merge request labels into a un-nested expr exposed field
	evalContext.MergeRequest = evalContext.Project.MergeRequest
	evalContext.Project.MergeRequest = nil

	evalContext.MergeRequest.Assignees = evalContext.MergeRequest.CurrentAssignees.Nodes
	evalContext.MergeRequest.CurrentAssignees = nil

	evalContext.MergeRequest.Reviewers = evalContext.MergeRequest.CurrentReviewers.Nodes
	evalContext.MergeRequest.CurrentReviewers = nil

	// Copy "current user" into MR
	evalContext.MergeRequest.CurrentUser = evalContext.CurrentUser

	evalContext.MergeRequest.Labels = evalContext.MergeRequest.ResponseLabels.Nodes
	evalContext.MergeRequest.ResponseLabels = nil

	evalContext.Group = evalContext.Project.ResponseGroup
	evalContext.Project.ResponseGroup = nil

	// Set top-level Pipeline: for pipeline webhook events fetch by ID in a separate low-complexity query;
	// for merge_request/note events use the MR's HeadPipeline.
	evalContext.Pipeline = evalContext.MergeRequest.HeadPipeline
	if pid := state.PipelineID(ctx); pid != "" {
		if p := fetchPipelineByID(ctx, client, state.ProjectID(ctx), pid); p != nil {
			evalContext.Pipeline = p
		}
	}
	if evalContext.Pipeline != nil && evalContext.Pipeline.ResponseJobs != nil {
		evalContext.Pipeline.Jobs = evalContext.Pipeline.ResponseJobs.Nodes
		evalContext.Pipeline.ResponseJobs = nil
	}

	evalContext.MergeRequest.Notes = evalContext.MergeRequest.ResponseNotes.Nodes
	evalContext.MergeRequest.ResponseNotes.Nodes = nil

	if evalContext.MergeRequest.ResponseFirstCommits != nil && len(evalContext.MergeRequest.ResponseFirstCommits.Nodes) > 0 {
		evalContext.MergeRequest.FirstCommit = &evalContext.MergeRequest.ResponseFirstCommits.Nodes[0]

		tmp := time.Since(*evalContext.MergeRequest.FirstCommit.CommittedDate)
		evalContext.MergeRequest.TimeSinceFirstCommit = &tmp
	}

	evalContext.MergeRequest.ResponseFirstCommits = nil

	if evalContext.MergeRequest.ResponseLastCommits != nil && len(evalContext.MergeRequest.ResponseLastCommits.Nodes) > 0 {
		evalContext.MergeRequest.LastCommit = &evalContext.MergeRequest.ResponseLastCommits.Nodes[0]

		tmp := time.Since(*evalContext.MergeRequest.LastCommit.CommittedDate)
		evalContext.MergeRequest.TimeSinceLastCommit = &tmp
	}

	evalContext.MergeRequest.ResponseLastCommits = nil

	if evalContext.MergeRequest.FirstCommit != nil && evalContext.MergeRequest.LastCommit != nil {
		tmp := evalContext.MergeRequest.FirstCommit.CommittedDate.Sub(*evalContext.MergeRequest.LastCommit.CommittedDate).Round(time.Hour)
		evalContext.MergeRequest.TimeBetweenFirstAndLastCommit = &tmp
	}

	return evalContext, nil
}

func (c *Context) IsValid() bool {
	return c != nil
}

func (c *Context) SetWebhookEvent(in any) {
	c.WebhookEvent = in
}

func (c *Context) SetContext(ctx context.Context) {
	c.Context = ctx
}

func (c *Context) GetDescription() string {
	if c.MergeRequest.Description == nil {
		return ""
	}

	return *c.MergeRequest.Description
}

func (c *Context) CanUseConfigurationFileFromChangeRequest(ctx context.Context) bool {
	// If the Merge Request has diverged from HEAD we can't trust the configuration
	if c.MergeRequest.DivergedFromTargetBranch {
		slogctx.Warn(ctx, "The Merge Request branch has diverged from HEAD; will use the scm-engine config from HEAD instead")

		return false
	}

	// If the Merge Request is not up to date with HEAD we can't trust the configuration
	if c.MergeRequest.ShouldBeRebased {
		slogctx.Warn(ctx, "The Merge Request branch is not up to date with HEAD; will use the scm-engine config from HEAD instead")

		return false
	}

	slogctx.Info(ctx, "The Merge Request branch is up to date with HEAD; will use the scm-engine config from the branch")

	return true
}

// AllowPipelineFailure controls if the CI pipeline are allowed to fail
//
// We allow the pipeline to fail with an error if the SCM-Engine configuration file
// is changed within the merge request, effectively allowing us to lint the configuration
// file when changing it but failing "open" in all other cases.
func (c *Context) AllowPipelineFailure(ctx context.Context) bool {
	return len(c.MergeRequest.findModifiedFiles(state.ConfigFilePath(ctx))) == 1
}

func (c *Context) TrackActionGroupExecution(group string) {
	// Ungrouped actions shouldn't be tracked
	if len(group) == 0 {
		return
	}

	c.ActionGroups[group] = true
}

func (c *Context) HasExecutedActionGroup(group string) bool {
	// Ungrouped actions shouldn't be tracked
	if len(group) == 0 {
		return false
	}

	_, ok := c.ActionGroups[group]

	return ok
}

// GetCodeOwners returns the eligible code owners for the merge request
//
// This is based on the elibible approvers in the rules of the merge request approval
// state api.
func (c *Context) GetCodeOwners() scm.Actors {
	actors := make(scm.Actors, 0)

	for _, rule := range c.MergeRequest.ApprovalState.Rules {
		// Multiple code owner paths could be matched when sections are used, so
		// flatten the list of eligible approvers
		if rule.Type == nil || *rule.Type != ApprovalRuleTypeCodeOwner {
			continue
		}

		// Note that the anyone who has authored a commit in the MR won't be considered
		// an eligible approver as part of the GitLab API response.
		if rule.EligibleApprovers != nil {
			for _, user := range rule.EligibleApprovers {
				if user.Bot == true {
					continue
				}

				actor := user.ToActor()
				if actors.Has(actor) {
					continue
				}

				actors.Add(actor)
			}
		}
	}

	return actors
}

func (c *Context) GetReviewers() scm.Actors {
	actors := make(scm.Actors, 0)

	for _, reviewer := range c.MergeRequest.Reviewers {
		actor := reviewer.ToActor()
		actors.Add(actor)
	}

	return actors
}

func (c *Context) GetAuthor() scm.Actor {
	return c.MergeRequest.Author.ToActor()
}

func (c *Context) GetLabels() []string {
	labels := make([]string, len(c.MergeRequest.Labels))
	for i, label := range c.MergeRequest.Labels {
		labels[i] = label.Title
	}

	return labels
}
