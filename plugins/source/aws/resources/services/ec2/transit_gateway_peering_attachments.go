// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TransitGatewayPeeringAttachments() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_transit_gateway_peering_attachments",
		Resolver:  fetchEc2TransitGatewayPeeringAttachments,
		Multiplex: client.ServiceAccountRegionMultiplexer("ec2"),
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
				Name:     "transit_gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "accepter_tgw_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccepterTgwInfo"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "requester_tgw_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RequesterTgwInfo"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "transit_gateway_attachment_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayAttachmentId"),
			},
		},
	}
}
