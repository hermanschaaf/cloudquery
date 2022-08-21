// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func ComputeSslPolicies() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_ssl_policies",
		Resolver:  fetchComputeSslPolicies,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "custom_features",
				Type: schema.TypeStringArray,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "enabled_features",
				Type: schema.TypeStringArray,
			},
			{
				Name: "fingerprint",
				Type: schema.TypeString,
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
				Name: "min_tls_version",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "profile",
				Type: schema.TypeString,
			},
			{
				Name: "warnings",
				Type: schema.TypeJSON,
			},
			{
				Name: "server_response",
				Type: schema.TypeJSON,
			},
		},
	}
}

func fetchComputeSslPolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.SslPolicies.List(c.ProjectId).PageToken(nextPageToken).Do()
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
