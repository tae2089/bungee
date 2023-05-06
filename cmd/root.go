package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
)

const (
	_defaultProfile = "default"
)

var (
	// rootCmd represents the base command when called without any sub-commands
	rootCmd = &cobra.Command{
		Use:   "bungee",
		Short: `Bungee is a CLI tool that is related to server connection. It supports AWS SSM connection, GCP IAP tunneling, SSH, and more.`,
		Long:  `Bungee is a CLI tool that is related to server connection. It supports AWS SSM connection, GCP IAP tunneling, SSH, and more.`,
	}

	_version                 string
	_credential              *Credential
	_credentialWithMFA       = fmt.Sprintf("%s_mfa", config.DefaultSharedCredentialsFilename())
	_credentialWithTemporary = fmt.Sprintf("%s_temporary", config.DefaultSharedCredentialsFilename())
)

type Credential struct {
	awsProfile    string
	awsConfig     *aws.Config
	gossmHomePath string
	ssmPluginPath string
}

func init() {
	// cobra.OnInitialize(initConfig)

}
