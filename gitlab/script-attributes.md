# Script Attributes

!!! tip "The [Expr Language Definition](https://expr-lang.org/docs/language-definition) is a great resource to learn more about the language"

!!! note

    Missing an attribute? The `schema/gitlab.schema.graphqls` file are what is used to query GitLab, adding the missing `field` to the right `type` should make it accessible. Please open an issue or Pull Request if something is missing.

The following attributes are available in `script` fields.

They can be accessed exactly as shown in this list.


## `project` {#project data-toc-label="project"}

The project the Merge Request belongs to

- `#!css project.archived` ; `boolean`. Indicates the archived status of the project
- `#!css project.created_at` ; `time`. Timestamp of the project creation
- `#!css project.description` ; `string`. Short description of the project
- `#!css project.full_path` ; `string`. Full path of the project
- `#!css project.id` ; `string`. ID of the project
- `#!css project.issues_enabled` ; `boolean`. Indicates if Issues are enabled for the current user
- `#!css project.last_activity_at` ; `time`. Timestamp of the project last activity
- `#!css project.name` ; `string`. Name of the project (without namespace)
- `#!css project.name_with_namespace` ; `string`. Full name of the project with its namespace
- `#!css project.path` ; `string`. Path of the project
- `#!css project.topics` ; `[]string`. List of project topics
- `#!css project.visibility` ; `string`. Visibility of the project

### `project.labels[]` {#project.labels[] data-toc-label="labels"}

Labels available on this project

- `#!css project.labels[].color` ; `string`. Background color of the label
- `#!css project.labels[].description` ; `string`. Description of the label (Markdown rendered as HTML for caching)
- `#!css project.labels[].id` ; `string`. Label ID
- `#!css project.labels[].title` ; `string`. Content of the label

## `group` {#group data-toc-label="group"}

The project group

- `#!css group.description` ; `string`. Description of the namespace
- `#!css group.emails_disabled` ; `optional boolean`. Indicates if a group has email notifications disabled
- `#!css group.full_name` ; `string`. Full name of the namespace
- `#!css group.full_path` ; `string`. Full path of the namespace
- `#!css group.id` ; `string`. ID of the namespace
- `#!css group.mentions_disabled` ; `optional boolean`. Indicates if a group is disabled from getting mentioned
- `#!css group.name` ; `string`. Name of the namespace
- `#!css group.path` ; `string`. Path of the namespace
- `#!css group.visibility` ; `optional string`. Visibility of the namespace
- `#!css group.web_url` ; `string`. Web URL of the group

## `merge_request` {#merge_request data-toc-label="merge_request"}

Information about the Merge Request

- `#!css merge_request.approvals_left` ; `optional int`. Number of approvals left
- `#!css merge_request.approvals_required` ; `optional int`. Number of approvals required
- `#!css merge_request.approved` ; `boolean`. Indicates if the merge request has all the required approvals
- `#!css merge_request.auto_merge_enabled` ; `boolean`. Indicates if auto merge is enabled for the merge request
- `#!css merge_request.auto_merge_strategy` ; `optional string`. Selected auto merge strategy
- `#!css merge_request.commit_count` ; `optional int`. Number of commits in the merge request
- `#!css merge_request.conflicts` ; `boolean`. Indicates if the merge request has conflicts
- `#!css merge_request.created_at` ; `time`. Timestamp of when the merge request was created
- `#!css merge_request.description` ; `optional string`. Description of the merge request (Markdown rendered as HTML for caching)
- `#!css merge_request.detailed_merge_status` (optional enum) Detailed merge status of the merge request

      *The following values are valid:*

      - `UNCHECKED` Merge status has not been checked
      - `CHECKING` Currently checking for mergeability
      - `MERGEABLE` Branch can be merged
      - `COMMITS_STATUS` Source branch exists and contains commits
      - `CI_MUST_PASS` Pipeline must succeed before merging
      - `CI_STILL_RUNNING` Pipeline is still running
      - `DISCUSSIONS_NOT_RESOLVED` Discussions must be resolved before merging
      - `DRAFT_STATUS` Merge request must not be draft before merging
      - `NOT_OPEN` Merge request must be open before merging
      - `NOT_APPROVED` Merge request must be approved before merging
      - `BLOCKED_STATUS` Merge request dependencies must be merged
      - `EXTERNAL_STATUS_CHECKS` Status checks must pass
      - `PREPARING` Merge request diff is being created
      - `JIRA_ASSOCIATION` Either the title or description must reference a Jira issue
      - `CONFLICT` There are conflicts between the source and target branches
      - `NEED_REBASE` Merge request needs to be rebased
      - `REQUESTED_CHANGES` Indicates a reviewer has requested change

- `#!css merge_request.discussion_locked` ; `boolean`. Indicates if comments on the merge request are locked to members only
- `#!css merge_request.diverged_from_target_branch` ; `boolean`. Indicates if the source branch is behind the target branch
- `#!css merge_request.downvotes` ; `int`. Number of downvotes for the merge request
- `#!css merge_request.draft` ; `boolean`. Indicates if the merge request is a draft
- `#!css merge_request.force_remove_source_branch` ; `optional boolean`. Indicates if the project settings will lead to source branch deletion after merge
- `#!css merge_request.id` ; `string`. ID of the merge request
- `#!css merge_request.iid` ; `string`. Internal ID of the merge request
- `#!css merge_request.merge_status_enum` (optional enum) Merge status of the merge request

      *The following values are valid:*

      - `UNCHECKED` Merge status has not been checked
      - `CHECKING` Currently checking for mergeability
      - `CAN_BE_MERGED` There are no conflicts between the source and target branches
      - `CANNOT_BE_MERGED` There are conflicts between the source and target branches
      - `CANNOT_BE_MERGED_RECHECK` Currently unchecked. The previous state was CANNOT_BE_MERGED

- `#!css merge_request.merge_when_pipeline_succeeds` ; `optional boolean`. Indicates if the merge has been set to auto-merge
- `#!css merge_request.mergeable` ; `boolean`. Indicates if the merge request is mergeable
- `#!css merge_request.mergeable_discussions_state` ; `optional boolean`. Indicates if all discussions in the merge request have been resolved, allowing the merge request to be merged
- `#!css merge_request.merged_at` ; `optional time`. Timestamp of when the merge request was merged, null if not merged
- `#!css merge_request.prepared_at` ; `optional time`. Timestamp of when the merge request was prepared
- `#!css merge_request.should_be_rebased` ; `boolean`. Indicates if the merge request will be rebased
- `#!css merge_request.should_remove_source_branch` ; `optional boolean`. Indicates if the source branch of the merge request will be deleted after merge
- `#!css merge_request.source_branch` ; `string`. Source branch of the merge request
- `#!css merge_request.source_branch_exists` ; `boolean`. Indicates if the source branch of the merge request exists
- `#!css merge_request.source_branch_protected` ; `boolean`. Indicates if the source branch is protected
- `#!css merge_request.squash` ; `boolean`. Indicates if the merge request is set to be squashed when merged. Project settings may override this value. Use squash_on_merge instead to take project squash options into account
- `#!css merge_request.squash_on_merge` ; `boolean`. Indicates if the merge request will be squashed when merged
- `#!css merge_request.state` ; `string`. State of the merge request
- `#!css merge_request.target_branch` ; `string`. Target branch of the merge request
- `#!css merge_request.target_branch_exists` ; `boolean`. Indicates if the target branch of the merge request exists
- `#!css merge_request.time_between_first_and_last_commit` ; `optional duration`. Duration between first and last commit made
- `#!css merge_request.time_since_first_commit` ; `optional duration`. Duration (from 'now') since the first commit was made
- `#!css merge_request.time_since_last_commit` ; `optional duration`. Duration (from 'now') since the last commit was made
- `#!css merge_request.title` ; `string`. Title of the merge request
- `#!css merge_request.updated_at` ; `time`. Timestamp of when the merge request was last updated
- `#!css merge_request.upvotes` ; `int`. Number of upvotes for the merge request.
- `#!css merge_request.user_discussions_count` ; `optional int`. Number of user discussions in the merge request
- `#!css merge_request.user_notes_count` ; `optional int`. User notes count of the merge request

### `merge_request.reviewers[]` {#merge_request.reviewers[] data-toc-label="reviewers"}

Users assigned to a merge request as a reviewer.

- `#!css merge_request.reviewers[].bot` ; `boolean`. Indicates if the user is a bot
- `#!css merge_request.reviewers[].id` ; `string`. ID of the user
- `#!css merge_request.reviewers[].public_email` ; `optional string`. User’s public email
- `#!css merge_request.reviewers[].state` (enum) State of the user

      *The following values are valid:*

      - `active` User is active and can use the system
      - `blocked` User has been blocked by an administrator and cannot use the system
      - `deactivated` User is no longer active and cannot use the system
      - `banned` User is blocked, and their contributions are hidden
      - `ldap_blocked` User has been blocked by the system
      - `blocked_pending_approval` User is blocked and pending approval

- `#!css merge_request.reviewers[].username` ; `string`. Username of the user. Unique within this instance of GitLab

### `merge_request.diff_stats[]` {#merge_request.diff_stats[] data-toc-label="diff_stats"}

Changes to a single file

- `#!css merge_request.diff_stats[].additions` ; `int`. Number of lines added to this file
- `#!css merge_request.diff_stats[].deletions` ; `int`. Number of lines deleted from this file
- `#!css merge_request.diff_stats[].path` ; `string`. File path, relative to repository root

### `merge_request.labels[]` {#merge_request.labels[] data-toc-label="labels"}

Labels available on this merge request

- `#!css merge_request.labels[].color` ; `string`. Background color of the label
- `#!css merge_request.labels[].description` ; `string`. Description of the label (Markdown rendered as HTML for caching)
- `#!css merge_request.labels[].id` ; `string`. Label ID
- `#!css merge_request.labels[].title` ; `string`. Content of the label

### `merge_request.head_pipeline` {#merge_request.head_pipeline data-toc-label="head_pipeline"}

Pipeline running on the branch HEAD of the merge request

- `#!css merge_request.head_pipeline.active` ; `boolean`. Indicates if the pipeline is active
- `#!css merge_request.head_pipeline.cancelable` ; `boolean`. Specifies if a pipeline can be canceled
- `#!css merge_request.head_pipeline.complete` ; `boolean`. Indicates if a pipeline is complete
- `#!css merge_request.head_pipeline.duration` ; `optional int`. Duration of the pipeline in seconds
- `#!css merge_request.head_pipeline.failure_reason` ; `optional string`. The reason why the pipeline failed
- `#!css merge_request.head_pipeline.finished_at` ; `optional time`. Timestamp of the pipeline's completion
- `#!css merge_request.head_pipeline.id` ; `string`. ID of the pipeline
- `#!css merge_request.head_pipeline.iid` ; `string`. Internal ID of the pipeline
- `#!css merge_request.head_pipeline.latest` ; `boolean`. If the pipeline is the latest one or not
- `#!css merge_request.head_pipeline.name` ; `optional string`. Name of the pipeline
- `#!css merge_request.head_pipeline.path` ; `optional string`. Relative path to the pipeline's page
- `#!css merge_request.head_pipeline.retryable` ; `boolean`. Specifies if a pipeline can be retried
- `#!css merge_request.head_pipeline.started_at` ; `optional time`. Timestamp when the pipeline was started
- `#!css merge_request.head_pipeline.status` (enum) Status of the pipeline

      *The following values are valid:*

      - `CANCELED` 
      - `CANCELING` 
      - `CREATED` 
      - `FAILED` 
      - `MANUAL` 
      - `PENDING` 
      - `PREPARING` 
      - `RUNNING` 
      - `SCHEDULED` 
      - `SKIPPED` 
      - `SUCCESS` 
      - `WAITING_FOR_CALLBACK` 
      - `WAITING_FOR_RESOURCE` 

- `#!css merge_request.head_pipeline.stuck` ; `boolean`. If the pipeline is stuck
- `#!css merge_request.head_pipeline.total_jobs` ; `int`. The total number of jobs in the pipeline
- `#!css merge_request.head_pipeline.updated_at` ; `time`. Timestamp of the pipeline's last activity
- `#!css merge_request.head_pipeline.warnings` ; `boolean`. Indicates if a pipeline has warnings

### `merge_request.notes[]` {#merge_request.notes[] data-toc-label="notes"}

All notes on this MR

- `#!css merge_request.notes[].body` ; `string`. Content of the note
- `#!css merge_request.notes[].created_at` ; `time`. Timestamp of the note creation
- `#!css merge_request.notes[].updated_at` ; `time`. Timestamp of the note’s last activity

#### `merge_request.notes[].author` {#merge_request.notes[].author data-toc-label="author"}

User who wrote the note

- `#!css merge_request.notes[].author.bot` ; `boolean`. Indicates if the user is a bot
- `#!css merge_request.notes[].author.id` ; `string`. ID of the user
- `#!css merge_request.notes[].author.public_email` ; `optional string`. User’s public email
- `#!css merge_request.notes[].author.state` (enum) State of the user

      *The following values are valid:*

      - `active` User is active and can use the system
      - `blocked` User has been blocked by an administrator and cannot use the system
      - `deactivated` User is no longer active and cannot use the system
      - `banned` User is blocked, and their contributions are hidden
      - `ldap_blocked` User has been blocked by the system
      - `blocked_pending_approval` User is blocked and pending approval

- `#!css merge_request.notes[].author.username` ; `string`. Username of the user. Unique within this instance of GitLab

### `merge_request.first_commit` {#merge_request.first_commit data-toc-label="first_commit"}

Information about the first commit made

- `#!css merge_request.first_commit.author_email` ; `optional string`. Commit author’s email
- `#!css merge_request.first_commit.author_name` ; `optional string`. Commit authors name
- `#!css merge_request.first_commit.authored_date` ; `optional time`. Timestamp of when the commit was authored
- `#!css merge_request.first_commit.committed_date` ; `optional time`. Timestamp of when the commit was committed
- `#!css merge_request.first_commit.committer_email` ; `optional string`. Email of the committer
- `#!css merge_request.first_commit.committer_name` ; `optional string`. Name of the committer
- `#!css merge_request.first_commit.description` ; `optional string`. Description of the commit message
- `#!css merge_request.first_commit.full_title` ; `optional string`. Full title of the commit message
- `#!css merge_request.first_commit.id` ; `optional string`. ID (global ID) of the commit
- `#!css merge_request.first_commit.message` ; `optional string`. Raw commit message
- `#!css merge_request.first_commit.sha` ; `string`. SHA1 ID of the commit
- `#!css merge_request.first_commit.short_id` ; `string`. Short SHA1 ID of the commit
- `#!css merge_request.first_commit.title` ; `optional string`. Title of the commit message
- `#!css merge_request.first_commit.web_url` ; `string`. Web URL of the commit

#### `merge_request.first_commit.author` {#merge_request.first_commit.author data-toc-label="author"}

Author of the commit

- `#!css merge_request.first_commit.author.bot` ; `boolean`. Indicates if the user is a bot
- `#!css merge_request.first_commit.author.id` ; `string`. ID of the user
- `#!css merge_request.first_commit.author.public_email` ; `optional string`. User’s public email
- `#!css merge_request.first_commit.author.state` (enum) State of the user

      *The following values are valid:*

      - `active` User is active and can use the system
      - `blocked` User has been blocked by an administrator and cannot use the system
      - `deactivated` User is no longer active and cannot use the system
      - `banned` User is blocked, and their contributions are hidden
      - `ldap_blocked` User has been blocked by the system
      - `blocked_pending_approval` User is blocked and pending approval

- `#!css merge_request.first_commit.author.username` ; `string`. Username of the user. Unique within this instance of GitLab

### `merge_request.last_commit` {#merge_request.last_commit data-toc-label="last_commit"}

Information about the last commit made

- `#!css merge_request.last_commit.author_email` ; `optional string`. Commit author’s email
- `#!css merge_request.last_commit.author_name` ; `optional string`. Commit authors name
- `#!css merge_request.last_commit.authored_date` ; `optional time`. Timestamp of when the commit was authored
- `#!css merge_request.last_commit.committed_date` ; `optional time`. Timestamp of when the commit was committed
- `#!css merge_request.last_commit.committer_email` ; `optional string`. Email of the committer
- `#!css merge_request.last_commit.committer_name` ; `optional string`. Name of the committer
- `#!css merge_request.last_commit.description` ; `optional string`. Description of the commit message
- `#!css merge_request.last_commit.full_title` ; `optional string`. Full title of the commit message
- `#!css merge_request.last_commit.id` ; `optional string`. ID (global ID) of the commit
- `#!css merge_request.last_commit.message` ; `optional string`. Raw commit message
- `#!css merge_request.last_commit.sha` ; `string`. SHA1 ID of the commit
- `#!css merge_request.last_commit.short_id` ; `string`. Short SHA1 ID of the commit
- `#!css merge_request.last_commit.title` ; `optional string`. Title of the commit message
- `#!css merge_request.last_commit.web_url` ; `string`. Web URL of the commit

#### `merge_request.last_commit.author` {#merge_request.last_commit.author data-toc-label="author"}

Author of the commit

- `#!css merge_request.last_commit.author.bot` ; `boolean`. Indicates if the user is a bot
- `#!css merge_request.last_commit.author.id` ; `string`. ID of the user
- `#!css merge_request.last_commit.author.public_email` ; `optional string`. User’s public email
- `#!css merge_request.last_commit.author.state` (enum) State of the user

      *The following values are valid:*

      - `active` User is active and can use the system
      - `blocked` User has been blocked by an administrator and cannot use the system
      - `deactivated` User is no longer active and cannot use the system
      - `banned` User is blocked, and their contributions are hidden
      - `ldap_blocked` User has been blocked by the system
      - `blocked_pending_approval` User is blocked and pending approval

- `#!css merge_request.last_commit.author.username` ; `string`. Username of the user. Unique within this instance of GitLab

### `merge_request.author` {#merge_request.author data-toc-label="author"}

User who created this merge request

- `#!css merge_request.author.bot` ; `boolean`. Indicates if the user is a bot
- `#!css merge_request.author.id` ; `string`. ID of the user
- `#!css merge_request.author.public_email` ; `optional string`. User’s public email
- `#!css merge_request.author.state` (enum) State of the user

      *The following values are valid:*

      - `active` User is active and can use the system
      - `blocked` User has been blocked by an administrator and cannot use the system
      - `deactivated` User is no longer active and cannot use the system
      - `banned` User is blocked, and their contributions are hidden
      - `ldap_blocked` User has been blocked by the system
      - `blocked_pending_approval` User is blocked and pending approval

- `#!css merge_request.author.username` ; `string`. Username of the user. Unique within this instance of GitLab

### `merge_request.assignees[]` {#merge_request.assignees[] data-toc-label="assignees"}

Users assigned to a merge request

- `#!css merge_request.assignees[].bot` ; `boolean`. Indicates if the user is a bot
- `#!css merge_request.assignees[].id` ; `string`. ID of the user
- `#!css merge_request.assignees[].public_email` ; `optional string`. User’s public email
- `#!css merge_request.assignees[].state` (enum) State of the user

      *The following values are valid:*

      - `active` User is active and can use the system
      - `blocked` User has been blocked by an administrator and cannot use the system
      - `deactivated` User is no longer active and cannot use the system
      - `banned` User is blocked, and their contributions are hidden
      - `ldap_blocked` User has been blocked by the system
      - `blocked_pending_approval` User is blocked and pending approval

- `#!css merge_request.assignees[].username` ; `string`. Username of the user. Unique within this instance of GitLab

### `merge_request.approval_state` {#merge_request.approval_state data-toc-label="approval_state"}

Information relating to rules that must be satisfied to merge this merge request


#### `merge_request.approval_state.rules[]` {#merge_request.approval_state.rules[] data-toc-label="rules"}

List of approval rules associated with the merge request

- `#!css merge_request.approval_state.rules[].name` ; `optional string`. Name of the rule
- `#!css merge_request.approval_state.rules[].type` (optional enum) Type of the rule

      *The following values are valid:*

      - `REGULAR` A regular approval rule
      - `CODE_OWNER` A code_owner approval rule
      - `REPORT_APPROVER` A report_approver approval rule
      - `ANY_APPROVER` A any_approver approval rule


##### `merge_request.approval_state.rules[].eligible_approvers[]` {#merge_request.approval_state.rules[].eligible_approvers[] data-toc-label="eligible_approvers"}

List of all users eligible to approve the merge request (defined explicitly and from associated groups)

- `#!css merge_request.approval_state.rules[].eligible_approvers[].bot` ; `boolean`. Indicates if the user is a bot
- `#!css merge_request.approval_state.rules[].eligible_approvers[].id` ; `string`. ID of the user
- `#!css merge_request.approval_state.rules[].eligible_approvers[].public_email` ; `optional string`. User’s public email
- `#!css merge_request.approval_state.rules[].eligible_approvers[].state` (enum) State of the user

      *The following values are valid:*

      - `active` User is active and can use the system
      - `blocked` User has been blocked by an administrator and cannot use the system
      - `deactivated` User is no longer active and cannot use the system
      - `banned` User is blocked, and their contributions are hidden
      - `ldap_blocked` User has been blocked by the system
      - `blocked_pending_approval` User is blocked and pending approval

- `#!css merge_request.approval_state.rules[].eligible_approvers[].username` ; `string`. Username of the user. Unique within this instance of GitLab

## `current_user` {#current_user data-toc-label="current_user"}

Get information about current user

- `#!css current_user.bot` ; `boolean`. Indicates if the user is a bot
- `#!css current_user.id` ; `string`. ID of the user
- `#!css current_user.public_email` ; `optional string`. User’s public email
- `#!css current_user.state` (enum) State of the user

      *The following values are valid:*

      - `active` User is active and can use the system
      - `blocked` User has been blocked by an administrator and cannot use the system
      - `deactivated` User is no longer active and cannot use the system
      - `banned` User is blocked, and their contributions are hidden
      - `ldap_blocked` User has been blocked by the system
      - `blocked_pending_approval` User is blocked and pending approval

- `#!css current_user.username` ; `string`. Username of the user. Unique within this instance of GitLab


## `webhook_event`

!!! tip "`webhook_event` attribute is only available in `server` mode"

    You have access to the raw webhook event payload via `webhook_event.*` attributes (not listed below) in Expr script fields when using [`server`](../commands/server.md) mode.

    See the [GitLab Webhook Events documentation](https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html) for available fields.

    The attributes are named _exactly_ as documented in the GitLab documentation.

- [`Comments`](https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html#comment-events) - A comment is made or edited on an issue or merge request.
- [`Merge request events`](https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html#merge-request-events) - A merge request is created, updated, or merged.
