package aws

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var (
	awsServiceSync sync.Once
	awsService     AwsServie
)

type AwsServie interface {
	GetEc2List() error
	StartInstance(instanceId []string) error
	StopInstances(instanceId []string) error
	ConnectInstances(instanceId string) error
}

func NewAwsService(client *ec2.Client, ssmClient *ssm.Client) AwsServie {
	if awsService == nil {
		awsServiceSync.Do(func() {
			awsService = &awsServiceImpl{
				client: client,
				ssm:    ssmClient,
			}
		})
	}
	return awsService
}
