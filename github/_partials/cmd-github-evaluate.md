NAME:
   scm-engine github evaluate - Evaluate a Pull Request

USAGE:
   scm-engine github evaluate [command options]  [pr_id, pr_id, ...]

OPTIONS:
   --project value  GitHub project (example: 'jippi/scm-engine') [$GITHUB_REPOSITORY]
   --id value       The Pull Request ID to process, if not provided as a CLI flag [$SCM_ENGINE_PULL_REQUEST_ID]
   --commit value   The git commit sha [$GITHUB_SHA]
   --help, -h       show help

GLOBAL OPTIONS:
   --config value  Path to the scm-engine config file (default: ".scm-engine.yml") [$SCM_ENGINE_CONFIG_FILE]
   --dry-run       Dry run, don't actually _do_ actions, just print them (default: false) [$SCM_ENGINE_DRY_RUN]
   --help, -h      show help
   --version, -v   print the version

