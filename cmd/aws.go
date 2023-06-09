package cmd

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// rootCmd represents the base command when called without any sub-commands
	awsCmd = &cobra.Command{
		Use:   "aws",
		Short: `The aws subcommand includes functionalities such as instance listing, SSM connection, and deletion.`,
		Long:  `The aws subcommand includes functionalities such as instance listing, SSM connection, and deletion.`,
	}
	_credential *Credential
)

type Credential struct {
	awsProfile    string
	awsConfig     *aws.Config
	gossmHomePath string
	ssmPluginPath string
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,	will be global for your application.
	// Cobra also supports local flags, which will only run when this action is called directly.
	ec2ListCmd.PersistentFlags().StringP("profile", "p", "", `[optional] if you are having multiple aws profiles, it is one of profiles (default is AWS_PROFILE environment variable or default)`)
	ec2ListCmd.PersistentFlags().StringP("region", "r", "", `[optional] it is region in AWS that would like to do something`)
	// mapping viper
	viper.BindPFlag("profile", ec2ListCmd.PersistentFlags().Lookup("profile"))
	viper.BindPFlag("region", ec2ListCmd.PersistentFlags().Lookup("region"))
	rootCmd.AddCommand(awsCmd)
}
