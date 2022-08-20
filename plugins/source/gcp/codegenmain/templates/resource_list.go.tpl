// Code generated by codegen; DO NOT EDIT.

package {{.GCPService}}

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func {{.GCPService | ToCamel}}{{.GCPSubService | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

func fetch{{.GCPService | ToCamel}}{{.GCPSubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.{{.GCPService | ToCamel}}.{{.GCPSubService | ToCamel}}.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
    res <- output.Items

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}