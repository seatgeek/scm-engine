package backstage_test

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/jippi/scm-engine/pkg/integration/backstage"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
	"gotest.tools/v3/assert"
)

func getRecorder(t *testing.T) *recorder.Recorder {
	t.Helper()

	fixtureName := "testdata/" + t.Name()

	hook := func(i *cassette.Interaction) error {
		if i.Request.Headers != nil && i.Request.Headers.Get("Authorization") != "" {
			i.Request.Headers.Set("Authorization", "REDACTED")
		}

		return nil
	}

	var opts []recorder.Option
	opts = []recorder.Option{
		recorder.WithRealTransport(&oauth2.Transport{
			Base: http.DefaultTransport,
			Source: oauth2.ReuseTokenSource(nil, oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: os.Getenv("BACKSTAGE_TOKEN"),
					TokenType:   "Bearer",
				},
			)),
		}),
		recorder.WithHook(hook, recorder.BeforeSaveHook),
		recorder.WithMatcher(cassette.MatcherFunc(func(r1 *http.Request, r2 cassette.Request) bool {
			// doesn't match automatically when providing real transport
			return r1.URL.String() == r2.URL
		})),
		recorder.WithMode(recorder.ModeRecordOnce),
	}

	r, err := recorder.New(fixtureName, opts...)
	if err != nil {
		t.Fatal(err)
	}

	return r
}

func TestClient_GetEntityOwner(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		filters []string
		want    *backstage.EntityReference
		wantErr error
	}{
		{
			name:    "not found",
			filters: []string{"kind=system,metadata.name=example"},
			want:    nil,
			wantErr: errors.New("No system found in Backstage catalog"),
		},
		{
			name:    "found",
			filters: []string{"kind=system,metadata.name=test-system"},
			want: &backstage.EntityReference{
				Kind:      "group",
				Namespace: "default",
				Name:      "test-group",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := getRecorder(t)
			defer r.Stop()

			client, err := backstage.NewClient(context.Background(), "https://backstage.example.com/", "", r.GetDefaultClient())
			require.NoError(t, err)

			entityRef, err := client.GetEntityOwner(context.Background(), tt.filters...)
			if tt.wantErr != nil {
				require.Error(t, err)
				assert.ErrorContains(t, tt.wantErr, err.Error())

				return
			}

			require.NoError(t, err)
			assert.DeepEqual(t, tt.want, entityRef)
		})
	}
}
