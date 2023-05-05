package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tae2089/bungee/internal/di"
)

var (
	// rootCmd represents the base command when called without any sub-commands
	ec2ListCmd = &cobra.Command{
		Use:   "list",
		Short: `This command invokes the functionality to retrieve a list of instances for AWS.`,
		Long: `
		This command provides the functionality to retrieve a list of instances for AWS, and also allows you to configure the region and profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			profile := strings.TrimSpace(viper.GetString("profile"))
			region := strings.TrimSpace(viper.GetString("region"))
			awsService := di.InitAwsService(profile, region)
			return awsService.GetEc2List()
			// awsService.ListInstances()
		},
	}
)

func init() {
	awsCmd.AddCommand(ec2ListCmd)
}
