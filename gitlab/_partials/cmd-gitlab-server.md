NAME:
   scm-engine gitlab server - Start HTTP server for webhook event driven usage

USAGE:
   scm-engine gitlab server [command options]

OPTIONS:
   --webhook-secret value                                                                           Used to validate received payloads. Sent with the request in the X-Gitlab-Token HTTP header [$SCM_ENGINE_WEBHOOK_SECRET]
   --listen-host value                                                                              IP that the HTTP server should listen on (default: "0.0.0.0") [$SCM_ENGINE_LISTEN_ADDR]
   --listen-port value                                                                              Port that the HTTP server should listen on (default: 3000) [$SCM_ENGINE_LISTEN_PORT, $PORT]
   --timeout value                                                                                  Timeout for webhook requests (default: 5s) [$SCM_ENGINE_TIMEOUT]
   --update-pipeline                                                                                Update the CI pipeline status with progress (default: true) [$SCM_ENGINE_UPDATE_PIPELINE]
   --update-pipeline-url value                                                                      (Optional) URL to where logs can be found for the pipeline [$SCM_ENGINE_UPDATE_PIPELINE_URL]
   --periodic-evaluation-interval value                                                             (Optional) Frequency of which to evaluate all Merge Requests regardless of user activity (default: 0s) [$SCM_ENGINE_PERIODIC_EVALUATION_INTERVAL]
   --periodic-evaluation-ignore-mr-labels value [ --periodic-evaluation-ignore-mr-labels value ]    (Optional) Ignore MR with these labels [$SCM_ENGINE_PERIODIC_EVALUATION_IGNORE_MR_WITH_LABELS]
   --periodic-evaluation-require-mr-labels value [ --periodic-evaluation-require-mr-labels value ]  (Optional) Only process MR with these labels [$SCM_ENGINE_PERIODIC_EVALUATION_REQUIRE_MR_WITH_LABELS]
   --periodic-evaluation-project-topics value [ --periodic-evaluation-project-topics value ]        (Optional) Only evaluate projects with these topics [$SCM_ENGINE_PERIODIC_EVALUATION_REQUIRE_PROJECT_TOPICS]
   --periodic-evaluation-only-project-membership                                                    (Optional) Only evaluate projects with membership (default: true) [$SCM_ENGINE_PERIODIC_EVALUATION_ONLY_PROJECTS_WITH_MEMBERSHIP]
   --help, -h                                                                                       show help

GLOBAL OPTIONS:
   --config value  Path to the scm-engine config file (default: ".scm-engine.yml") [$SCM_ENGINE_CONFIG_FILE]
   --dry-run       Dry run, don't actually _do_ actions, just print them (default: false) [$SCM_ENGINE_DRY_RUN]
   --help, -h      show help
   --version, -v   print the version

