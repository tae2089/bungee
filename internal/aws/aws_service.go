package aws

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var (
	awsServiceSync sync.Once
	awsService     AwsServie
)

type AwsServie interface {
	GetEc2List() error
}

func NewAwsService(client *ec2.Client) AwsServie {
	if awsService == nil {
		awsServiceSync.Do(func() {
			awsService = &awsServiceImpl{
				client: client,
			}
		})
	}
	return awsService
}
