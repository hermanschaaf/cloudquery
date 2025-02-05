package cloudtrail

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type CloudTrailWrapper struct {
	types.Trail
	Tags map[string]string
}

// groupNameRegex extracts log group name from the ARN
var groupNameRegex = regexp.MustCompile("arn:[a-zA-Z0-9-]+:logs:[a-z0-9-]+:[0-9]+:log-group:([a-zA-Z0-9-/]+):")

func fetchCloudtrailTrails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Cloudtrail
	log := meta.(*client.Client).Logger()
	response, err := svc.DescribeTrails(ctx, nil)

	if err != nil {
		return err
	}

	getBundledTrailsWithTags := func(trails []types.Trail, region string) ([]CloudTrailWrapper, error) {
		processed := make([]CloudTrailWrapper, len(trails))

		input := cloudtrail.ListTagsInput{
			ResourceIdList: make([]string, 0, len(trails)),
		}

		for i, h := range trails {
			processed[i] = CloudTrailWrapper{
				Trail: h,
				Tags:  make(map[string]string),
			}

			// Before fetching trail tags we have to check if the trail is organization trail
			// If the trail is organization trail and the account id is not matched with current account id
			// We skip, and not fetch the trail tags
			arnParts, err := arn.Parse(*h.TrailARN)
			if err != nil {
				log.Warn().Str("arn", *h.TrailARN).Msg("cloud not parse cloudtrail ARN")
				continue
			}
			if aws.ToBool(h.IsOrganizationTrail) && c.AccountID != arnParts.AccountID {
				log.Warn().Str("arn", *h.TrailARN).Msg("the trail is an organization level trail, cloud not fetch tags")
				continue
			}

			input.ResourceIdList = append(input.ResourceIdList, *h.TrailARN)
		}

		if len(input.ResourceIdList) == 0 {
			return processed, nil
		}

		for {
			response, err := svc.ListTags(ctx, &input, func(options *cloudtrail.Options) {
				options.Region = region
			})
			if err != nil {
				return nil, err
			}
			for i, tr := range processed {
				client.TagsIntoMap(getCloudTrailTagsByResourceID(*tr.TrailARN, response.ResourceTagList), processed[i].Tags)
			}
			if aws.ToString(response.NextToken) == "" {
				break
			}
			input.NextToken = response.NextToken
		}

		return processed, nil
	}

	// since api returns all the cloudtrails despite region we aggregate trails by region to get tags.
	aggregatedTrails, err := aggregateCloudTrails(response.TrailList)
	if err != nil {
		return err
	}
	for region, trails := range aggregatedTrails {
		for i := 0; i < len(trails); i += 20 {
			end := i + 20

			if end > len(trails) {
				end = len(trails)
			}
			t := trails[i:end]
			processed, err := getBundledTrailsWithTags(t, region)
			if err != nil {
				return err
			}
			res <- processed
		}
	}

	return nil
}

func postCloudtrailTrailResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Cloudtrail
	r := resource.Item.(CloudTrailWrapper)
	response, err := svc.GetTrailStatus(ctx,
		&cloudtrail.GetTrailStatusInput{Name: r.TrailARN}, func(o *cloudtrail.Options) {
			o.Region = *r.HomeRegion
		})
	if err != nil {
		return err
	}
	if err := resource.Set("is_logging", response.IsLogging); err != nil {
		return err
	}
	if err := resource.Set("latest_cloud_watch_logs_delivery_error", response.LatestCloudWatchLogsDeliveryError); err != nil {
		return err
	}
	if err := resource.Set("latest_cloud_watch_logs_delivery_time", response.LatestCloudWatchLogsDeliveryTime); err != nil {
		return err
	}
	if err := resource.Set("latest_delivery_error", response.LatestDeliveryError); err != nil {
		return err
	}
	if err := resource.Set("latest_delivery_time", response.LatestDeliveryTime); err != nil {
		return err
	}
	if err := resource.Set("latest_digest_delivery_error", response.LatestDigestDeliveryError); err != nil {
		return err
	}
	if err := resource.Set("latest_digest_delivery_time", response.LatestDigestDeliveryTime); err != nil {
		return err
	}
	if err := resource.Set("latest_notification_error", response.LatestNotificationError); err != nil {
		return err
	}
	if err := resource.Set("latest_notification_time", response.LatestNotificationTime); err != nil {
		return err
	}
	if err := resource.Set("start_logging_time", response.StartLoggingTime); err != nil {
		return err
	}
	return resource.Set("stop_logging_time", response.StopLoggingTime)
}

func resolveCloudtrailTrailCloudwatchLogsLogGroupName(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	groupName := ""
	log := meta.(*client.Client).Logger()
	r := resource.Item.(CloudTrailWrapper)
	if r.CloudWatchLogsLogGroupArn != nil {
		matches := groupNameRegex.FindStringSubmatch(*r.CloudWatchLogsLogGroupArn)
		if len(matches) < 2 {
			log.Warn().Str("arn", *r.CloudWatchLogsLogGroupArn).Msg("CloudWatchLogsLogGroupARN doesn't fit standard regex")
		} else {
			groupName = matches[1]
		}
	} else {
		log.Info().Msg("CloudWatchLogsLogGroupARN is empty")
	}

	return resource.Set("cloudwatch_logs_log_group_name", groupName)
}

func fetchCloudtrailTrailEventSelectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(CloudTrailWrapper)
	c := meta.(*client.Client)
	svc := c.Services().Cloudtrail
	response, err := svc.GetEventSelectors(ctx, &cloudtrail.GetEventSelectorsInput{TrailName: r.TrailARN}, func(options *cloudtrail.Options) {
		options.Region = *r.HomeRegion
	})
	if err != nil {
		return err
	}
	res <- response.EventSelectors
	return nil
}

func getCloudTrailTagsByResourceID(id string, set []types.ResourceTag) []types.Tag {
	for _, s := range set {
		if *s.ResourceId == id {
			return s.TagsList
		}
	}
	return nil
}

func aggregateCloudTrails(trails []types.Trail) (map[string][]types.Trail, error) {
	resp := make(map[string][]types.Trail)
	for _, t := range trails {
		if t.HomeRegion == nil {
			return nil, fmt.Errorf("got cloudtrail with HomeRegion == nil")
		}
		resp[*t.HomeRegion] = append(resp[*t.HomeRegion], t)
	}
	return resp, nil
}
