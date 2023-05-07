package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tae2089/bungee/internal/di"
)

var (
	// rootCmd represents the base command when called without any sub-commands
	instanceId  = ""
	ec2StartCmd = &cobra.Command{
		Use:   "start",
		Short: `The aws subcommand includes functionalities such as instance listing, SSM connection, and deletion.`,
		Long:  `The aws subcommand includes functionalities such as instance listing, SSM connection, and deletion.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			profile := strings.TrimSpace(viper.GetString("profile"))
			region := strings.TrimSpace(viper.GetString("region"))
			awsService := di.InitAwsService(profile, region)
			return awsService.StartInstance(args)
			// awsService.ListInstances()
		},
	}
)

func init() {
	awsCmd.AddCommand(ec2StartCmd)
}
