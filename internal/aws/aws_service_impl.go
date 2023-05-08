package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/olekukonko/tablewriter"
)

var _ AwsServie = (*awsServiceImpl)(nil)

type awsServiceImpl struct {
	client *ec2.Client
}

// StartInstance implements AwsServie
func (a *awsServiceImpl) StartInstance(instanceId []string) error {
	input := &ec2.StartInstancesInput{
		InstanceIds: instanceId,
	}
	// 인스턴스 시작 요청 전송
	_, err := a.client.StartInstances(context.TODO(), input)
	if err != nil {
		return err
	}
	fmt.Println("Instance started:", instanceId)
	return nil
}

// StopInstances implements AwsServie
func (a *awsServiceImpl) StopInstances(instanceId []string) error {
	input := &ec2.StopInstancesInput{
		InstanceIds: instanceId,
	}

	// 인스턴스 stop 요청 전송
	output, err := a.client.StopInstances(context.TODO(), input)

	if err != nil {
		return err
	}

	for _, instance := range output.StoppingInstances {
		fmt.Println("Instance stoped:", *instance.InstanceId)
	}

	return nil
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
