package di

import (
	"fmt"

	"github.com/tae2089/bungee/internal/aws"
	"github.com/tae2089/bungee/internal/config"
	"github.com/tae2089/bungee/internal/ssh"
)

func InitAwsService(profile, region string) aws.AwsServie {
	client, err := config.GetEc2Client(profile, region)
	if err != nil {
		return nil
	}
	ssmClient, err := config.GetSsmClient(profile, region)
	if err != nil {
		return nil
	}
	awsService := aws.NewAwsService(client, ssmClient)
	return awsService
}

func InitSshService() ssh.SSHService {
	db, err := config.NewDbConfig()
	if err != nil {
		fmt.Println(err)
		db.Close()
		return nil
	}
	sshService := ssh.NewSshService(db)

	return sshService
}
