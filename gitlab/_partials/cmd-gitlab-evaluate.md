NAME:
   scm-engine gitlab evaluate - Evaluate a Merge Request

USAGE:
   scm-engine gitlab evaluate [command options]  [mr_id, mr_id, ...]

OPTIONS:
   --update-pipeline            Update the CI pipeline status with progress (default: true) [$SCM_ENGINE_UPDATE_PIPELINE]
   --update-pipeline-url value  (Optional) URL to where logs can be found for the pipeline [$SCM_ENGINE_UPDATE_PIPELINE_URL]
   --project value              GitLab project (example: 'gitlab-org/gitlab') [$GITLAB_PROJECT, $CI_PROJECT_PATH]
   --id value                   The Merge Request ID to process, if not provided as a CLI flag [$CI_MERGE_REQUEST_IID]
   --commit value               The git commit sha [$CI_COMMIT_SHA]
   --help, -h                   show help

GLOBAL OPTIONS:
   --config value  Path to the scm-engine config file (default: ".scm-engine.yml") [$SCM_ENGINE_CONFIG_FILE]
   --dry-run       Dry run, don't actually _do_ actions, just print them (default: false) [$SCM_ENGINE_DRY_RUN]
   --help, -h      show help
   --version, -v   print the version

