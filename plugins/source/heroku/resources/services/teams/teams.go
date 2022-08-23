package teams

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource teams --config gen.hcl --output .
func Teams() *schema.Table {
	return &schema.Table{
		Name:        "github_teams",
		Description: "Teams allow you to manage access to a shared group of applications and other resources",
		Resolver:    fetchTeams,
		IgnoreError: client.IgnoreError,
		Columns: []schema.Column{
			{
				Name:        "created_at",
				Description: "when the team was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "credit_card_collections",
				Description: "whether charges incurred by the team are paid by credit card",
				Type:        schema.TypeBool,
			},
			{
				Name:        "default",
				Description: "whether to use this team when none is specified",
				Type:        schema.TypeBool,
			},
			{
				Name:     "enterprise_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EnterpriseAccount.ID"),
			},
			{
				Name:     "enterprise_account_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EnterpriseAccount.Name"),
			},
			{
				Name:        "id",
				Description: "unique identifier of team",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:     "identity_provider_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityProvider.ID"),
			},
			{
				Name:     "identity_provider_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityProvider.Name"),
			},
			{
				Name:     "identity_provider_owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityProvider.Owner.ID"),
			},
			{
				Name:     "identity_provider_owner_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityProvider.Owner.Name"),
			},
			{
				Name:     "identity_provider_owner_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityProvider.Owner.Type"),
			},
			{
				Name:        "membership_limit",
				Description: "upper limit of members allowed in a team",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "name",
				Description: "unique name of team",
				Type:        schema.TypeString,
			},
			{
				Name:        "provisioned_licenses",
				Description: "whether the team is provisioned licenses by salesforce",
				Type:        schema.TypeBool,
			},
			{
				Name:        "role",
				Description: "role in the team",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "type of team",
				Type:        schema.TypeString,
			},
			{
				Name:        "updated_at",
				Description: "when the team was updated",
				Type:        schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchTeams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
