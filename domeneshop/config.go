package domeneshop

import (
	"context"
	"fmt"

	"github.com/MindTooth/go-domeneshop"
)

type Config struct {
	Token            string
	Secret           string
	TerraformVersion string
}

type Client struct {
	client *domeneshop.APIClient
	config *Config
	ctx    *context.Context
}

func (c *Client) getClient() *domeneshop.APIClient {
	return c.client

}

func (c *Client) getContext() *context.Context {
	return c.ctx
}

func (c *Config) Client() (*Client, error) {
	cfg := domeneshop.NewConfiguration()

	cfg.UserAgent = fmt.Sprintf("Terraform/%s", c.TerraformVersion)

	doClient := domeneshop.NewAPIClient(cfg)

	ctx := context.WithValue(context.Background(), domeneshop.ContextBasicAuth, domeneshop.BasicAuth{
		UserName: c.Token,
		Password: c.Secret,
	})

	provider := &Client{
		client: doClient,
		config: c,
		ctx:    &ctx,
	}

	return provider, nil
}
