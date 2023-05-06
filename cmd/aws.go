package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// rootCmd represents the base command when called without any sub-commands
	awsCmd = &cobra.Command{
		Use:   "aws",
		Short: `The aws subcommand includes functionalities such as instance listing, SSM connection, and deletion.`,
		Long:  `The aws subcommand includes functionalities such as instance listing, SSM connection, and deletion.`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)

func init() {

	rootCmd.AddCommand(awsCmd)
}
