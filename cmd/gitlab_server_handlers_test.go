package cmd_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jippi/scm-engine/cmd"
	"github.com/jippi/scm-engine/pkg/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGitLabWebhookHandler_PipelinePayloadWithNilMergeRequest_Returns400(t *testing.T) {
	t.Parallel()

	// Pipeline payload with object_attributes set but merge_request omitted (nil).
	// GitLab sends this for branch/tag pipelines that are not associated with an MR.
	body := `{
		"object_kind": "pipeline",
		"project": {
			"path_with_namespace": "group/my-project"
		},
		"object_attributes": {
			"iid": 12345,
			"commit": {"id": "abc123def"},
			"last_commit": {}
		}
	}`

	ctx := t.Context()
	ctx = state.WithProvider(ctx, "gitlab")
	ctx = state.WithToken(ctx, "fake-token")
	ctx = state.WithBaseURL(ctx, "http://fake-gitlab.example")
	ctx = state.WithBackstageURL(ctx, "")
	ctx = state.WithBackstageToken(ctx, "")

	handler := cmd.GitLabWebhookHandler(ctx, "")

	req := httptest.NewRequest(http.MethodPost, "/gitlab", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code, "expected 400 when pipeline payload has no merge_request")
	assert.Contains(t, rec.Body.String(), "not associated with a merge request",
		"response body should explain that payload is not associated with a merge request")
}
