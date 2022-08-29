// Code generated by codegen; DO NOT EDIT.

package run

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:      "gcp_run_services",
		Resolver:  fetchServices,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "api_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiVersion"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "spec",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
		},
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Run.Projects.Locations.Services.List("projects/" + c.ProjectId + "/locations/-").Continue(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- output.Items
		if output.Metadata == nil || output.Metadata.Continue == "" {
			break
		}
		nextPageToken = output.Metadata.Continue
	}
	return nil
}
