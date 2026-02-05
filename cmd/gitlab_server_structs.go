package cmd

type GitlabWebhookPayload struct {
	EventType        string                                `json:"event_type"`
	Project          GitlabWebhookPayloadProject           `json:"project"`                     // "project" is sent for all events
	ObjectAttributes *GitlabWebhookPayloadObjectAttributes `json:"object_attributes,omitempty"` // "object_attributes" is sent on "merge_request" events and "pipeline" events
	MergeRequest     *GitlabWebhookPayloadMergeRequest     `json:"merge_request,omitempty"`     // "merge_request" is sent on "note" activity and "pipeline" events
}

type GitlabWebhookPayloadProject struct {
	PathWithNamespace string `json:"path_with_namespace"`
}

type GitlabWebhookPayloadObjectAttributes struct {
	IID        int                        `json:"iid"`
	LastCommit GitlabWebhookPayloadCommit `json:"last_commit"`
	Commit     GitlabWebhookPayloadCommit `json:"commit"`
}

func (o *GitlabWebhookPayloadObjectAttributes) GetCommitID() string {
	if o.LastCommit != nil {
		return o.LastCommit.ID
	}
	return o.Commit.ID
}

type GitlabWebhookPayloadMergeRequest struct {
	IID        int                        `json:"iid"`
	LastCommit GitlabWebhookPayloadCommit `json:"last_commit"`
}

type GitlabWebhookPayloadCommit struct {
	ID string `json:"id"`
}
