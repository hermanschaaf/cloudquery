// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:      "heroku_team",
		Resolver:  fetchTeams,
		Multiplex: client.NoMultiplex,
		Columns: []schema.Column{
			{
				Name: "created_at",
				Type: schema.TypeJSON,
			},
			{
				Name: "credit_card_collections",
				Type: schema.TypeBool,
			},
			{
				Name: "default",
				Type: schema.TypeBool,
			},
			{
				Name: "enterprise_account",
				Type: schema.TypeJSON,
			},
			{
				Name: "id",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "identity_provider",
				Type: schema.TypeJSON,
			},
			{
				Name: "membership_limit",
				Type: schema.TypeFloat,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "provisioned_licenses",
				Type: schema.TypeBool,
			},
			{
				Name: "role",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "updated_at",
				Type: schema.TypeJSON,
			},
		},
	}
}

func fetchTeams(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	v, err := c.Services.Heroku.TeamList(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- v
	return nil
}
