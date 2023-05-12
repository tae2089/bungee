package cmd

import (
	"github.com/spf13/cobra"
)

// create ssh command using cobra
var (
	// rootCmd represents the base command when called without any sub-commands
	sshCmd = &cobra.Command{
		Use:   "ssh",
		Short: `The aws subcommand includes functionalities such as instance listing, SSM connection, and deletion.`,
		Long:  `The aws subcommand includes functionalities such as instance listing, SSM connection, and deletion.`,
	}
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,	will be global for your application.
	rootCmd.AddCommand(sshCmd)
}
