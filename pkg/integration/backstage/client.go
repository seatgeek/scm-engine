package backstage

import (
	"context"
	"net/http"

	go_backstage "github.com/datolabs-io/go-backstage/v3"
	"golang.org/x/oauth2"
)

var _ Client = (*client)(nil)

type Client interface {
	ListEntities(ctx context.Context, options *go_backstage.ListEntityOptions) ([]go_backstage.Entity, error)
}

type client struct {
	wrapped *go_backstage.Client
}

// NewClient creates a new Backstage client with an optional bearer token
func NewClient(ctx context.Context, baseURL string, token string) (Client, error) {
	var httpClient *http.Client

	if token != "" {
		httpClient = oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: token,
			TokenType:   "Bearer",
		}))
	}

	backstageClient, err := go_backstage.NewClient(baseURL, token, httpClient)
	if err != nil {
		return nil, err
	}

	return &client{wrapped: backstageClient}, nil
}

func (c *client) ListEntities(ctx context.Context, options *go_backstage.ListEntityOptions) ([]go_backstage.Entity, error) {
	entities, response, err := c.wrapped.Catalog.Entities.List(ctx, options)
	response.Body.Close()

	if err != nil {
		return nil, err
	}

	return entities, nil
}
