package di

import (
	"github.com/tae2089/bungee/internal/aws"
	"github.com/tae2089/bungee/internal/config"
)

func InitAwsService(profile, region string) aws.AwsServie {
	client, err := config.GetEc2Client(profile, region)
	if err != nil {
		return nil
	}
	awsService := aws.NewAwsService(client)
	return awsService
}
