// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"

	"google.golang.org/api/compute/v1"
)

func ComputeUrlMaps() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_url_maps",
		Resolver:  fetchComputeUrlMaps,
		Multiplex: client.ProjectMultiplex,
		Options: schema.TableCreationOptions{
			PrimaryKeys: []string{
				"project_id",
				"id",
			},
		},
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "default_route_action",
				Type: schema.TypeJSON,
			},
			{
				Name: "default_service",
				Type: schema.TypeString,
			},
			{
				Name: "default_url_redirect",
				Type: schema.TypeJSON,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "fingerprint",
				Type: schema.TypeString,
			},
			{
				Name: "header_action",
				Type: schema.TypeJSON,
			},
			{
				Name: "host_rules",
				Type: schema.TypeJSON,
			},
			{
				Name: "id",
				Type: schema.TypeInt,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "path_matchers",
				Type: schema.TypeJSON,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "tests",
				Type: schema.TypeJSON,
			},
		},
	}
}

func fetchComputeUrlMaps(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.UrlMaps.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		var allItems []*compute.UrlMap
		for _, items := range output.Items {
			allItems = append(allItems, items.UrlMaps...)
		}
		res <- allItems

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
