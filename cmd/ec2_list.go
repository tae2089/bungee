package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tae2089/bungee/internal/di"
)

var (
	// rootCmd represents the base command when called without any sub-commands
	ec2ListCmd = &cobra.Command{
		Use:   "list",
		Short: `The aws subcommand includes functionalities such as instance listing, SSM connection, and deletion.`,
		Long:  `The aws subcommand includes functionalities such as instance listing, SSM connection, and deletion.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			awsService := di.InitAwsService("", "")
			return awsService.GetEc2List()
			// awsService.ListInstances()
		},
	}
)

func init() {

	awsCmd.AddCommand(ec2ListCmd)
}
