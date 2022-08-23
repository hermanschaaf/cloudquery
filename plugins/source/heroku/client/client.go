package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	heroku "github.com/heroku/heroku-go/v5"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger
	Heroku HerokuServices
	Team   string
	Teams  []string
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func (c Client) WithTeam(team string) schema.ClientMeta {
	return &Client{
		logger: c.logger.With("team", team),
		Heroku: c.Heroku,
		Team:   team,
		Teams:  c.Teams,
	}
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	providerConfig := config.(*Config)

	// TODO: validate provider config
	heroku.DefaultTransport.BearerToken = providerConfig.Token
	h := heroku.NewService(heroku.DefaultClient)

	return &Client{
		logger: logger,
		Heroku: HerokuServices{
			Teams: h,
		},
		Teams: providerConfig.Teams,
	}, nil
}
