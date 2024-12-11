package backstage

import (
	"context"
	"errors"
	"net/http"

	go_backstage "github.com/datolabs-io/go-backstage/v3"
	"golang.org/x/oauth2"
)

type Client struct {
	wrapped *go_backstage.Client
}

// NewClient creates a new Backstage client with an optional bearer token
func NewClient(ctx context.Context, baseURL string, token string, httpClient *http.Client) (*Client, error) {
	if token != "" && httpClient == nil {
		httpClient = oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: token,
			TokenType:   "Bearer",
		}))
	}

	if baseURL == "" {
		return nil, errors.New("baseURL is required")
	}

	backstageClient, err := go_backstage.NewClient(baseURL, "default", httpClient)
	if err != nil {
		return nil, err
	}

	return &Client{wrapped: backstageClient}, nil
}

func (c *Client) ListEntities(ctx context.Context, options *go_backstage.ListEntityOptions) ([]go_backstage.Entity, error) {
	entities, response, err := c.wrapped.Catalog.Entities.List(ctx, options)
	response.Body.Close()

	if err != nil {
		return nil, err
	}

	return entities, nil
}
