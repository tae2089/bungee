package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/olekukonko/tablewriter"
)

type awsServiceImpl struct {
	client *ec2.Client
}

// GetEc2List implements AwsServie
func (a *awsServiceImpl) GetEc2List() error {
	input := &ec2.DescribeInstancesInput{}
	result, err := a.client.DescribeInstances(context.Background(), input)
	if err != nil {
		fmt.Println("failed to describe instances:", err)
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Instance ID", "Instance Type", "Availability Zone", "Public IP", "Private IP", "Key Name", "State"})

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			table.Append([]string{
				*instance.InstanceId,
				aws.ToString((*string)(&instance.InstanceType)),
				*instance.Placement.AvailabilityZone,
				aws.ToString(instance.PublicIpAddress),
				*instance.PrivateIpAddress,
				aws.ToString(instance.KeyName),
				aws.ToString((*string)(&instance.State.Name)),
				// *instance.KeyName,
			})
		}
	}

	table.Render()

	return nil
}

var _ AwsServie = (*awsServiceImpl)(nil)
