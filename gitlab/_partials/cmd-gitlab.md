NAME:
   scm-engine gitlab - GitLab related commands

USAGE:
   scm-engine gitlab command [command options]

COMMANDS:
   lint      lint a configuration file
   evaluate  Evaluate a Merge Request
   help, h   Shows a list of commands or help for one command

OPTIONS:
   --api-token value  GitLab API token [$SCM_ENGINE_TOKEN]
   --base-url value   Base URL for the SCM instance (default: "https://gitlab.com/") [$SCM_ENGINE_BASE_URL, $CI_SERVER_URL]
   --help, -h         show help

GLOBAL OPTIONS:
   --config value  Path to the scm-engine config file (default: ".scm-engine.yml") [$SCM_ENGINE_CONFIG_FILE]
   --dry-run       Dry run, don't actually _do_ actions, just print them (default: false) [$SCM_ENGINE_DRY_RUN]
   --help, -h      show help
   --version, -v   print the version

