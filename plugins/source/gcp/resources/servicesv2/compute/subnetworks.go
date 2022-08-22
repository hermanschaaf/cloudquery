// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"

	"google.golang.org/api/compute/v1"
)

func ComputeSubnetworks() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_subnetworks",
		Resolver:  fetchComputeSubnetworks,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "enable_flow_logs",
				Type: schema.TypeBool,
			},
			{
				Name: "external_ipv_6_prefix",
				Type: schema.TypeString,
			},
			{
				Name: "fingerprint",
				Type: schema.TypeString,
			},
			{
				Name: "gateway_address",
				Type: schema.TypeString,
			},
			{
				Name: "id",
				Type: schema.TypeInt,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "internal_ipv_6_prefix",
				Type: schema.TypeString,
			},
			{
				Name: "ip_cidr_range",
				Type: schema.TypeString,
			},
			{
				Name: "ipv_6_access_type",
				Type: schema.TypeString,
			},
			{
				Name: "ipv_6_cidr_range",
				Type: schema.TypeString,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "log_config",
				Type: schema.TypeJSON,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "network",
				Type: schema.TypeString,
			},
			{
				Name: "private_ip_google_access",
				Type: schema.TypeBool,
			},
			{
				Name: "private_ipv_6_google_access",
				Type: schema.TypeString,
			},
			{
				Name: "purpose",
				Type: schema.TypeString,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "role",
				Type: schema.TypeString,
			},
			{
				Name: "secondary_ip_ranges",
				Type: schema.TypeJSON,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "stack_type",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
		},
	}
}

func fetchComputeSubnetworks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.Subnetworks.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		var allItems []*compute.Subnetwork
		for _, items := range output.Items {
			allItems = append(allItems, items.Subnetworks...)
		}
		res <- allItems

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
