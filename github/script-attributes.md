# Script Attributes

!!! tip "The [Expr Language Definition](https://expr-lang.org/docs/language-definition) is a great resource to learn more about the language"

!!! note

    Missing an attribute? The `schema/gitlab.schema.graphqls` file are what is used to query GitLab, adding the missing `field` to the right `type` should make it accessible. Please open an issue or Pull Request if something is missing.

The following attributes are available in `script` fields.

They can be accessed exactly as shown in this list.


## `repository` {#repository data-toc-label="repository"}

The project the Pull Request belongs to

- `#!css repository.allow_update_branch` ; `boolean`. Whether or not a Pull Request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging
- `#!css repository.archived_at` ; `optional time`. Identifies the date and time when the repository was archived
- `#!css repository.auto_merge_allowed` ; `boolean`. Whether or not Auto-merge can be enabled on Pull Requests in this repository
- `#!css repository.created_at` ; `time`. Identifies the date and time when the object was created
- `#!css repository.delete_branch_on_merge` ; `boolean`. Whether or not branches are automatically deleted when merged in this repository
- `#!css repository.description` ; `optional string`. The description of the repository
- `#!css repository.has_discussions_enabled` ; `boolean`. Indicates if the repository has the Discussions feature enabled
- `#!css repository.has_issues_enabled` ; `boolean`. Indicates if the repository has issues feature enabled
- `#!css repository.has_projects_enabled` ; `boolean`. Indicates if the repository has the Projects feature enabled
- `#!css repository.has_wiki_enabled` ; `boolean`. Indicates if the repository has wiki feature enabled
- `#!css repository.id` ; `string`. The Node ID of the Repository object
- `#!css repository.is_archived` ; `boolean`. Indicates if the repository is unmaintained
- `#!css repository.is_blank_issues_enabled` ; `boolean`. Returns true if blank issue creation is allowed
- `#!css repository.is_disabled` ; `boolean`. Returns whether or not this repository disabled
- `#!css repository.is_fork` ; `boolean`. Identifies if the repository is a fork
- `#!css repository.is_locked` ; `boolean`. Indicates if the repository has been locked or not
- `#!css repository.is_mirror` ; `boolean`. Identifies if the repository is a mirror
- `#!css repository.is_private` ; `boolean`. Identifies if the repository is private or internal
- `#!css repository.is_template` ; `boolean`. Identifies if the repository is a template that can be used to generate new repositories
- `#!css repository.is_user_configuration_repository` ; `boolean`. Is this repository a user configuration repository
- `#!css repository.merge_commit_allowed` ; `boolean`. Whether or not PRs are merged with a merge commit on this repository
- `#!css repository.name` ; `string`. The name of the repository
- `#!css repository.name_with_owner` ; `string`. The repository's name with owner
- `#!css repository.pushed_at` ; `optional time`. Identifies the date and time when the repository was last pushed to
- `#!css repository.rebase_merge_allowed` ; `boolean`. Whether or not rebase-merging is enabled on this repository
- `#!css repository.resource_path` ; `string`. The HTTP path for this repository
- `#!css repository.squash_merge_allowed` ; `boolean`. Whether or not squash-merging is enabled on this repository
- `#!css repository.stargazer_count` ; `int`. Returns a count of how many stargazers there are on this object
- `#!css repository.updated_at` ; `time`. Identifies the date and time when the object was last updated
- `#!css repository.url` ; `string`. The HTTP URL for this repository
- `#!css repository.visibility` (enum) Indicates the repository's visibility level

      *The following values are valid:*

      - `PRIVATE` The repository is visible only to those with explicit access
      - `PUBLIC` The repository is visible to everyone
      - `INTERNAL` The repository is visible only to users in the same business


## `owner` {#owner data-toc-label="owner"}

The project owner

- `#!css owner.login` ; `string`. The username used to login

## `pull_request` {#pull_request data-toc-label="pull_request"}

Information about the Pull Request

- `#!css pull_request.active_lock_reason` (optional enum) Reason that the conversation was locked

      *The following values are valid:*

      - `OFF_TOPIC` The issue or Pull Request was locked because the conversation was off-topic
      - `TOO_HEATED` The issue or Pull Request was locked because the conversation was too heated
      - `RESOLVED` The issue or Pull Request was locked because the conversation was resolved
      - `SPAM` The issue or Pull Request was locked because the conversation was spam

- `#!css pull_request.additions` ; `int`. The number of additions in this Pull Request
- `#!css pull_request.base_ref_name` ; `string`. Identifies the name of the base Ref associated with the Pull Request, even if the ref has been deleted
- `#!css pull_request.body` ; `string`. The body as Markdown
- `#!css pull_request.can_be_rebased` ; `boolean`. Whether or not the Pull Request is rebaseable
- `#!css pull_request.changed_files` ; `int`. The number of changed files in this Pull Request
- `#!css pull_request.closed` ; `boolean`. `true` if the Pull Request is closed
- `#!css pull_request.closed_at` ; `optional time`. Identifies the date and time when the object was closed
- `#!css pull_request.created_at` ; `time`. Identifies the date and time when the object was created
- `#!css pull_request.created_via_email` ; `boolean`. Check if this comment was created via an email reply
- `#!css pull_request.deletions` ; `int`. The number of deletions in this Pull Request
- `#!css pull_request.head_ref_name` ; `string`. Identifies the name of the head Ref associated with the Pull Request, even if the ref has been deleted
- `#!css pull_request.id` ; `string`. The Node ID of the PullRequest object
- `#!css pull_request.includes_created_edit` ; `boolean`. Check if this comment was edited and includes an edit with the creation data
- `#!css pull_request.is_cross_repository` ; `boolean`. The head and base repositories are different
- `#!css pull_request.is_draft` ; `boolean`. Identifies if the Pull Request is a draft
- `#!css pull_request.is_in_merge_queue` ; `boolean`. Indicates whether the Pull Request is in a merge queue
- `#!css pull_request.is_merge_queue_enabled` ; `boolean`. Indicates whether the Pull Request's base ref has a merge queue enabled
- `#!css pull_request.last_edited_at` ; `optional time`. The moment the editor made the last edit
- `#!css pull_request.locked` ; `boolean`. `true` if the Pull Request is locked
- `#!css pull_request.maintainer_can_modify` ; `boolean`. Indicates whether maintainers can modify the Pull Request
- `#!css pull_request.merge_state_status` (enum) Detailed information about the current Pull Request merge state status

      *The following values are valid:*

      - `DIRTY` The merge commit cannot be cleanly created
      - `UNKNOWN` The state cannot currently be determined
      - `BLOCKED` The merge is blocked
      - `BEHIND` The head ref is out of date
      - `UNSTABLE` Mergeable with non-passing commit status
      - `HAS_HOOKS` Mergeable with passing commit status and pre-receive hooks
      - `CLEAN` Mergeable and passing commit status

- `#!css pull_request.mergeable` (enum) Whether or not the Pull Request can be merged based on the existence of merge conflicts

      *The following values are valid:*

      - `MERGEABLE` The Pull Request can be merged
      - `CONFLICTING` The Pull Request cannot be merged due to merge conflicts
      - `UNKNOWN` The mergeability of the Pull Request is still being calculated

- `#!css pull_request.merged` ; `boolean`. Whether or not the Pull Request was merged
- `#!css pull_request.merged_at` ; `optional time`. The date and time that the Pull Request was merged
- `#!css pull_request.number` ; `int`. Identifies the Pull Request number
- `#!css pull_request.permalink` ; `string`. The permalink to the Pull Request
- `#!css pull_request.published_at` ; `optional time`. Identifies when the comment was published at
- `#!css pull_request.resource_path` ; `string`. The HTTP path for this Pull Request
- `#!css pull_request.review_decision` (enum) The current status of this Pull Request with respect to code review

      *The following values are valid:*

      - `CHANGES_REQUESTED` Changes have been requested on the Pull Request
      - `APPROVED` The Pull Request has received an approving review
      - `REVIEW_REQUIRED` A review is required before the Pull Request can be merged

- `#!css pull_request.state` (enum) Identifies the state of the Pull Request

      *The following values are valid:*

      - `OPEN` A Pull Request that is still open
      - `CLOSED` A Pull Request that has been closed without being merged
      - `MERGED` A Pull Request that has been closed by being merged

- `#!css pull_request.time_between_first_and_last_commit` ; `optional duration`. Duration between first and last commit made
- `#!css pull_request.time_since_first_commit` ; `optional duration`. Duration (from 'now') since the first commit was made
- `#!css pull_request.time_since_last_commit` ; `optional duration`. Duration (from 'now') since the last commit was made
- `#!css pull_request.title` ; `string`. Identifies the Pull Request title
- `#!css pull_request.total_comments_count` ; `optional int`. Returns a count of how many comments this Pull Request has received
- `#!css pull_request.updated_at` ; `time`. Identifies the date and time when the object was last updated
- `#!css pull_request.url` ; `string`. The HTTP URL for this Pull Request

### `pull_request.files[]` {#pull_request.files[] data-toc-label="files"}

- `#!css pull_request.files[].additions` ; `int`. The number of additions to the file
- `#!css pull_request.files[].change_type` (enum) How the file was changed in this PullRequest

      *The following values are valid:*

      - `ADDED` The file was added. Git status 'A'
      - `DELETED` The file was deleted. Git status 'D'
      - `RENAMED` The file was renamed. Git status 'R'
      - `COPIED` The file was copied. Git status 'C'
      - `MODIFIED` The file's contents were changed. Git status 'M'
      - `CHANGED` The file's type was changed. Git status 'T'

- `#!css pull_request.files[].deletions` ; `int`. The number of deletions to the file
- `#!css pull_request.files[].path` ; `string`. The path of the file

### `pull_request.first_commit` {#pull_request.first_commit data-toc-label="first_commit"}

Information about the first commit made

- `#!css pull_request.first_commit.additions` ; `int`. The number of additions in this commit
- `#!css pull_request.first_commit.authored_by_committer` ; `boolean`. Check if the committer and the author match
- `#!css pull_request.first_commit.authored_date` ; `time`. The datetime when this commit was authored
- `#!css pull_request.first_commit.changed_files_if_available` ; `optional int`. The number of changed files in this commit. If GitHub is unable to calculate the number of changed files (for example due to a timeout), this will return null. We recommend using this field instead of changedFiles
- `#!css pull_request.first_commit.commit_resource_path` ; `string`. The HTTP path for this Git object
- `#!css pull_request.first_commit.commit_url` ; `string`. The HTTP URL for this Git object
- `#!css pull_request.first_commit.committed_date` ; `time`. The datetime when this commit was committed
- `#!css pull_request.first_commit.committed_via_web` ; `boolean`. Check if committed via GitHub web UI
- `#!css pull_request.first_commit.deletions` ; `int`. The number of deletions in this commit
- `#!css pull_request.first_commit.message` ; `string`. The Git commit message
- `#!css pull_request.first_commit.message_body` ; `string`. The Git commit message body
- `#!css pull_request.first_commit.message_headline` ; `string`. The Git commit message headline
- `#!css pull_request.first_commit.url` ; `string`. The HTTP URL for this commit

#### `pull_request.first_commit.author` {#pull_request.first_commit.author data-toc-label="author"}

Authorship details of the commit

- `#!css pull_request.first_commit.author.date` ; `optional time`. The timestamp of the Git action (authoring or committing)
- `#!css pull_request.first_commit.author.email` ; `optional string`. The email in the Git commit
- `#!css pull_request.first_commit.author.name` ; `optional string`. The name in the Git commit

##### `pull_request.first_commit.author.user` {#pull_request.first_commit.author.user data-toc-label="user"}

The GitHub user corresponding to the email field. Null if no such user exists

- `#!css pull_request.first_commit.author.user.login` ; `string`. The username used to login

#### `pull_request.first_commit.committer` {#pull_request.first_commit.committer data-toc-label="committer"}

Committer details of the commit

- `#!css pull_request.first_commit.committer.date` ; `optional time`. The timestamp of the Git action (authoring or committing)
- `#!css pull_request.first_commit.committer.email` ; `optional string`. The email in the Git commit
- `#!css pull_request.first_commit.committer.name` ; `optional string`. The name in the Git commit

##### `pull_request.first_commit.committer.user` {#pull_request.first_commit.committer.user data-toc-label="user"}

The GitHub user corresponding to the email field. Null if no such user exists

- `#!css pull_request.first_commit.committer.user.login` ; `string`. The username used to login

### `pull_request.last_commit` {#pull_request.last_commit data-toc-label="last_commit"}

Information about the last commit made

- `#!css pull_request.last_commit.additions` ; `int`. The number of additions in this commit
- `#!css pull_request.last_commit.authored_by_committer` ; `boolean`. Check if the committer and the author match
- `#!css pull_request.last_commit.authored_date` ; `time`. The datetime when this commit was authored
- `#!css pull_request.last_commit.changed_files_if_available` ; `optional int`. The number of changed files in this commit. If GitHub is unable to calculate the number of changed files (for example due to a timeout), this will return null. We recommend using this field instead of changedFiles
- `#!css pull_request.last_commit.commit_resource_path` ; `string`. The HTTP path for this Git object
- `#!css pull_request.last_commit.commit_url` ; `string`. The HTTP URL for this Git object
- `#!css pull_request.last_commit.committed_date` ; `time`. The datetime when this commit was committed
- `#!css pull_request.last_commit.committed_via_web` ; `boolean`. Check if committed via GitHub web UI
- `#!css pull_request.last_commit.deletions` ; `int`. The number of deletions in this commit
- `#!css pull_request.last_commit.message` ; `string`. The Git commit message
- `#!css pull_request.last_commit.message_body` ; `string`. The Git commit message body
- `#!css pull_request.last_commit.message_headline` ; `string`. The Git commit message headline
- `#!css pull_request.last_commit.url` ; `string`. The HTTP URL for this commit

#### `pull_request.last_commit.author` {#pull_request.last_commit.author data-toc-label="author"}

Authorship details of the commit

- `#!css pull_request.last_commit.author.date` ; `optional time`. The timestamp of the Git action (authoring or committing)
- `#!css pull_request.last_commit.author.email` ; `optional string`. The email in the Git commit
- `#!css pull_request.last_commit.author.name` ; `optional string`. The name in the Git commit

##### `pull_request.last_commit.author.user` {#pull_request.last_commit.author.user data-toc-label="user"}

The GitHub user corresponding to the email field. Null if no such user exists

- `#!css pull_request.last_commit.author.user.login` ; `string`. The username used to login

#### `pull_request.last_commit.committer` {#pull_request.last_commit.committer data-toc-label="committer"}

Committer details of the commit

- `#!css pull_request.last_commit.committer.date` ; `optional time`. The timestamp of the Git action (authoring or committing)
- `#!css pull_request.last_commit.committer.email` ; `optional string`. The email in the Git commit
- `#!css pull_request.last_commit.committer.name` ; `optional string`. The name in the Git commit

##### `pull_request.last_commit.committer.user` {#pull_request.last_commit.committer.user data-toc-label="user"}

The GitHub user corresponding to the email field. Null if no such user exists

- `#!css pull_request.last_commit.committer.user.login` ; `string`. The username used to login

### `pull_request.merged_by` {#pull_request.merged_by data-toc-label="merged_by"}

The actor who merged the Pull Request

- `#!css pull_request.merged_by.login` ; `string`. The username used to login

### `pull_request.author` {#pull_request.author data-toc-label="author"}

The actor who authored the comment

- `#!css pull_request.author.login` ; `string`. The username used to login

### `pull_request.labels[]` {#pull_request.labels[] data-toc-label="labels"}

Labels available on this project

- `#!css pull_request.labels[].color` ; `string`. Identifies the label color
- `#!css pull_request.labels[].created_at` ; `optional time`. Identifies the date and time when the label was created
- `#!css pull_request.labels[].description` ; `optional string`. A brief description of this label
- `#!css pull_request.labels[].id` ; `string`. The Node ID of the Label object
- `#!css pull_request.labels[].is_default` ; `boolean`. Indicates whether or not this is a default label
- `#!css pull_request.labels[].name` ; `string`. Identifies the label name
- `#!css pull_request.labels[].updated_at` ; `time`. Identifies the date and time when the label was last updated

## `viewer` {#viewer data-toc-label="viewer"}

Get information about current user

- `#!css viewer.login` ; `string`. The username used to login


## `webhook_event`

!!! tip "`webhook_event` attribute is only available in `server` mode"

    You have access to the raw webhook event payload via `webhook_event.*` attributes (not listed below) in Expr script fields when using [`server`](../commands/server.md) mode.

    See the [GitLab Webhook Events documentation](https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html) for available fields.

    The attributes are named _exactly_ as documented in the GitLab documentation.

- [`Comments`](https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html#comment-events) - A comment is made or edited on an issue or merge request.
- [`Merge request events`](https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html#merge-request-events) - A merge request is created, updated, or merged.
