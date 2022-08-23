package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func HerokuMockTestHelper(t *testing.T, table *schema.Table, createService func() (*Services, error), options TestOptions) {
	t.Helper()

	table.IgnoreInTests = false

	newTestExecutionClient := func(ctx context.Context, p *plugins.SourcePlugin, spec specs.Source) (schema.ClientMeta, error) {
		svc, err := createService()
		if err != nil {
			return nil, fmt.Errorf("failed to create service %w", err)
		}
		var uSpec Spec
		if err := spec.UnmarshalSpec(&uSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
		}
		c := &Client{
			plugin:   p,
			Services: svc,
			projects: []string{"testProject"},
		}

		return c, nil
	}

	p := plugins.NewSourcePlugin(
		table.Name,
		"dev",
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:   "dev",
		Tables: []string{"*"},
	})
}
