package client

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin   *plugins.SourcePlugin
	projects []string
	backoff  BackoffSettings
	// All heroku services initialized by client
	Services *Services
	// this is set by table client multiplexer
	ProjectId string
	// Logger
	logger zerolog.Logger
}

const (
	defaultProjectIdName = "<CHANGE_THIS_TO_YOUR_PROJECT_ID>"
	serviceAccountEnvKey = "CQ_SERVICE_ACCOUNT_KEY_JSON"
)

//revive:disable:modifies-value-receiver

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func Configure(ctx context.Context, p *plugins.SourcePlugin, s specs.Source) (schema.ClientMeta, error) {
	c := Client{
		plugin: p,
	}
	var herokuSpec Spec
	if err := s.UnmarshalSpec(&herokuSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
	}

	return &c, nil
}
