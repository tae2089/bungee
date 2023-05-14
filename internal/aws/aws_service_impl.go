package aws

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/smithy-go"
	"github.com/olekukonko/tablewriter"
)

var _ AwsServie = (*awsServiceImpl)(nil)

type awsServiceImpl struct {
	client *ec2.Client
	ssm    *ssm.Client
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

// ConnectInstances implements AwsServie
func (a *awsServiceImpl) ConnectInstances(instanceId string) error {
	// AWS 세션 생성

	// Session Manager를 통해 EC2 인스턴스에 연결
	// 세션 시작 요청
	input := &ssm.StartSessionInput{
		Target: aws.String(instanceId),
	}
	startSessionOutput, err := a.ssm.StartSession(context.TODO(), input)
	// sessJson, err := json.Marshal(startSessionOutput)
	// if err != nil {
	// 	return err
	// }

	// paramsJson, err := json.Marshal(input)
	// if err != nil {
	// 	return err
	// }
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) {
			if apiErr.ErrorCode() == "TargetNotConnected" {
				return fmt.Errorf("failed to start session: target is not connected or does not exist")
			}
		}
		return err
	}

	// 세션 ID와 세션 토큰 얻기
	sessionID := *startSessionOutput.SessionId
	// sessionToken := *startSessionOutput.Token

	// 세션 터미널 입출력 처리
	err = callProcess("")
	if err != nil {
		return err
	}

	// 세션 종료 요청
	_, err = a.ssm.TerminateSession(context.TODO(), &ssm.TerminateSessionInput{
		SessionId: aws.String(sessionID),
	})
	if err != nil {
		return err
	}

	fmt.Println("Session terminated")

	return nil
}

// CallProcess calls process.
func callProcess(process string, args ...string) error {
	call := exec.Command(process, args...)
	call.Stderr = os.Stderr
	call.Stdout = os.Stdout
	call.Stdin = os.Stdin

	// ignore signal(sigint)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	done := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-sigs:
			case <-done:
				break
			}
		}
	}()
	defer close(done)

	// run subprocess
	if err := call.Run(); err != nil {
		return err
	}
	return nil
}
