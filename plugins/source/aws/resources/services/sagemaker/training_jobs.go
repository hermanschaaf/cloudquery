// Code generated by codegen; DO NOT EDIT.

package sagemaker

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TrainingJobs() *schema.Table {
	return &schema.Table{
		Name:      "aws_sagemaker_training_jobs",
		Resolver:  fetchSagemakerTrainingJobs,
		Multiplex: client.ServiceAccountRegionMultiplexer("api.sagemaker"),
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
				Resolver: schema.PathResolver("TrainingJobArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerTrainingJobTags,
				Description: `The tags associated with the model.`,
			},
			{
				Name:     "algorithm_specification",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AlgorithmSpecification"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "model_artifacts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ModelArtifacts"),
			},
			{
				Name:     "resource_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourceConfig"),
			},
			{
				Name:     "secondary_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecondaryStatus"),
			},
			{
				Name:     "stopping_condition",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StoppingCondition"),
			},
			{
				Name:     "training_job_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TrainingJobName"),
			},
			{
				Name:     "training_job_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TrainingJobStatus"),
			},
			{
				Name:     "auto_ml_job_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoMLJobArn"),
			},
			{
				Name:     "billable_time_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BillableTimeInSeconds"),
			},
			{
				Name:     "checkpoint_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CheckpointConfig"),
			},
			{
				Name:     "debug_hook_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DebugHookConfig"),
			},
			{
				Name:     "debug_rule_configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DebugRuleConfigurations"),
			},
			{
				Name:     "debug_rule_evaluation_statuses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DebugRuleEvaluationStatuses"),
			},
			{
				Name:     "enable_inter_container_traffic_encryption",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableInterContainerTrafficEncryption"),
			},
			{
				Name:     "enable_managed_spot_training",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableManagedSpotTraining"),
			},
			{
				Name:     "enable_network_isolation",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableNetworkIsolation"),
			},
			{
				Name:     "environment",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Environment"),
			},
			{
				Name:     "experiment_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExperimentConfig"),
			},
			{
				Name:     "failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailureReason"),
			},
			{
				Name:     "final_metric_data_list",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FinalMetricDataList"),
			},
			{
				Name:     "hyper_parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HyperParameters"),
			},
			{
				Name:     "input_data_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InputDataConfig"),
			},
			{
				Name:     "labeling_job_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LabelingJobArn"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "output_data_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OutputDataConfig"),
			},
			{
				Name:     "profiler_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProfilerConfig"),
			},
			{
				Name:     "profiler_rule_configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProfilerRuleConfigurations"),
			},
			{
				Name:     "profiler_rule_evaluation_statuses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProfilerRuleEvaluationStatuses"),
			},
			{
				Name:     "profiling_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProfilingStatus"),
			},
			{
				Name:     "retry_strategy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RetryStrategy"),
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleArn"),
			},
			{
				Name:     "secondary_status_transitions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecondaryStatusTransitions"),
			},
			{
				Name:     "tensor_board_output_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TensorBoardOutputConfig"),
			},
			{
				Name:     "training_end_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("TrainingEndTime"),
			},
			{
				Name:     "training_start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("TrainingStartTime"),
			},
			{
				Name:     "training_time_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TrainingTimeInSeconds"),
			},
			{
				Name:     "tuning_job_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TuningJobArn"),
			},
			{
				Name:     "vpc_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcConfig"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
