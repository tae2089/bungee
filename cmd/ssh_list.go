package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tae2089/bungee/internal/di"
)

var (
	// rootCmd represents the base command when called without any sub-commands
	sshListCmd = &cobra.Command{
		Use:   "list",
		Short: `This command invokes the functionality to retrieve a list of instances for AWS.`,
		Long: `
		This command provides the functionality to retrieve a list of instances for AWS, and also allows you to configure the region and profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			sshService := di.InitSshService()
			if sshService == nil {
				return fmt.Errorf("Failed to initialize ssh service")
			}
			err := sshService.ShowSSHInfoList()
			return err
		},
		SilenceUsage: true,
	}
)

func init() {
	sshCmd.AddCommand(sshListCmd)
}
