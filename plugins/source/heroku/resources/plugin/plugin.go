package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

const exampleConfig = `
kind: source
spec:
  tables: ["*"]
  spec:
    token: <YOUR_HEROKU_API_TOKEN>
    # Optional. GRPC Retry/backoff configuration, time units in seconds. Documented in https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md
    # backoff_base_delay: 1
    # backoff_multiplier: 1.6
    # backoff_max_delay: 120
    # backoff_jitter: 0.2
    # backoff_min_connect_timeout = 0
    # Optional. Max amount of retries for retrier, defaults to max 3 retries.
    # max_retries: 3
`

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"heroku",
		Version,
		[]*schema.Table{},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
		plugins.WithClassifyError(client.ClassifyError),
	)
}
