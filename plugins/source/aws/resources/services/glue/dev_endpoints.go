// Code generated by codegen; DO NOT EDIT.

package glue

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DevEndpoints() *schema.Table {
	return &schema.Table{
		Name:      "aws_glue_dev_endpoints",
		Resolver:  fetchGlueDevEndpoints,
		Multiplex: client.ServiceAccountRegionMultiplexer("glue"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGlueDevEndpointArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueDevEndpointTags,
			},
			{
				Name:     "arguments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Arguments"),
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "created_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTimestamp"),
			},
			{
				Name:     "endpoint_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EndpointName"),
			},
			{
				Name:     "extra_jars_s_3_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExtraJarsS3Path"),
			},
			{
				Name:     "extra_python_libs_s_3_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExtraPythonLibsS3Path"),
			},
			{
				Name:     "failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailureReason"),
			},
			{
				Name:     "glue_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlueVersion"),
			},
			{
				Name:     "last_modified_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedTimestamp"),
			},
			{
				Name:     "last_update_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdateStatus"),
			},
			{
				Name:     "number_of_nodes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumberOfNodes"),
			},
			{
				Name:     "number_of_workers",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumberOfWorkers"),
			},
			{
				Name:     "private_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateAddress"),
			},
			{
				Name:     "public_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicAddress"),
			},
			{
				Name:     "public_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicKey"),
			},
			{
				Name:     "public_keys",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PublicKeys"),
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleArn"),
			},
			{
				Name:     "security_configuration",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityConfiguration"),
			},
			{
				Name:     "security_group_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SecurityGroupIds"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "subnet_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetId"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
			{
				Name:     "worker_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkerType"),
			},
			{
				Name:     "yarn_endpoint_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("YarnEndpointAddress"),
			},
			{
				Name:     "zeppelin_remote_spark_interpreter_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ZeppelinRemoteSparkInterpreterPort"),
			},
		},
	}
}
