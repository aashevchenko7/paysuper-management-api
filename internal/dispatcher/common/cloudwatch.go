package common

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type CloudWatchInterface interface {
	FilterLogEventsWithContext(ctx aws.Context, input *cloudwatchlogs.FilterLogEventsInput,
		opts ...request.Option) (*cloudwatchlogs.FilterLogEventsOutput, error)
}

type CloudWatch struct {
	cloudWatchLogs *cloudwatchlogs.CloudWatchLogs
}

func NewCloudWatch(cfg *LogsSettings) (CloudWatchInterface, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(cfg.AwsCloudWatchRegion),
			Credentials: credentials.NewStaticCredentials(
				cfg.AwsCloudWatchAccessKeyId,
				cfg.AwsCloudWatchSecretAccessKey,
				"",
			),
		},
	)

	if err != nil {
		return nil, err
	}

	return &CloudWatch{cloudWatchLogs: cloudwatchlogs.New(sess)}, nil
}

func (m *CloudWatch) FilterLogEventsWithContext(
	ctx aws.Context,
	input *cloudwatchlogs.FilterLogEventsInput,
	opts ...request.Option,
) (*cloudwatchlogs.FilterLogEventsOutput, error) {
	return m.cloudWatchLogs.FilterLogEventsWithContext(ctx, input, opts...)
}
