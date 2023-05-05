package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func getAwsConfig(profile, region string) (aws.Config, error) {
	if profile == "" {
		profile = "default"
	}
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region))
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}

func GetEc2Client(profile, region string) (*ec2.Client, error) {
	cfg, err := getAwsConfig(profile, region)
	if err != nil {
		return nil, err
	}
	return ec2.NewFromConfig(cfg), nil
}
