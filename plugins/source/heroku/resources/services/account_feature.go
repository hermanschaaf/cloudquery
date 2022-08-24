// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func AccountFeatures() *schema.Table {
	return &schema.Table{
		Name:      "heroku_account_features",
		Resolver:  fetchAccountFeatures,
		Multiplex: client.NoMultiplex,
		Columns: []schema.Column{
			{
				Name: "created_at",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "display_name",
				Type: schema.TypeString,
			},
			{
				Name:     "doc_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DocURL"),
			},
			{
				Name: "enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "feedback_email",
				Type: schema.TypeString,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name: "updated_at",
				Type: schema.TypeTimestamp,
			},
		},
	}
}

func fetchAccountFeatures(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	v, err := c.Heroku.AccountFeatureList(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- v
	return nil
}
