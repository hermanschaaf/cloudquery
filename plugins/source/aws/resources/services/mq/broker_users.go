// Code generated by codegen; DO NOT EDIT.

package mq

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BrokerUsers() *schema.Table {
	return &schema.Table{
		Name:      "aws_mq_broker_users",
		Resolver:  fetchMqBrokerUsers,
		Multiplex: client.ServiceAccountRegionMultiplexer("mq"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "broker_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "broker_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BrokerId"),
			},
			{
				Name:     "console_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ConsoleAccess"),
			},
			{
				Name:     "groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Groups"),
			},
			{
				Name:     "pending",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Pending"),
			},
			{
				Name:     "username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Username"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
