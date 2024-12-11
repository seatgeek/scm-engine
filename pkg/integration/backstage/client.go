package backstage

import (
	"context"
	"net/http"

	go_backstage "github.com/datolabs-io/go-backstage/v3"
	"golang.org/x/oauth2"
)

type Client struct {
	wrapped *go_backstage.Client
}

// NewClient creates a new Backstage client with an optional bearer token
func NewClient(ctx context.Context, baseURL string, token string) (*Client, error) {
	var httpClient *http.Client

	if token != "" {
		httpClient = oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: token,
			TokenType:   "Bearer",
		}))
	}

	client, err := go_backstage.NewClient(baseURL, token, httpClient)
	if err != nil {
		return nil, err
	}

	return &Client{wrapped: client}, nil
}
