package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tae2089/bungee/internal/di"
	"github.com/tae2089/bungee/internal/ssh"
)

var (
	// rootCmd represents the base command when called without any sub-commands
	sshRegisterDto = &ssh.SshRegisterDto{}
	sshRegisterCmd = &cobra.Command{
		Use:   "register",
		Short: `This command invokes the functionality to retrieve a list of instances for AWS.`,
		Long: `
		This command provides the functionality to retrieve a list of instances for AWS, and also allows you to configure the region and profile.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// variables in ssh.sshRegisterDto  not nil or empty
			if sshRegisterDto.Host == "" {
				return fmt.Errorf("Host flag is required")
			}
			if sshRegisterDto.Name == "" {
				return fmt.Errorf("Name flag is required")
			}
			if sshRegisterDto.User == "" {
				return fmt.Errorf("User flag is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			sshService := di.InitSshService()
			if sshService == nil {
				return fmt.Errorf("Failed to initialize ssh service")
			}
			err := sshService.RegisterSSHInfo(sshRegisterDto)
			return err
		},
		SilenceUsage: true,
	}
)

func init() {
	sshRegisterCmd.Flags().StringVarP(&sshRegisterDto.Host, "host", "", "", "The host to connect to")
	sshRegisterCmd.Flags().StringVarP(&sshRegisterDto.Name, "Name", "n", "", "The name of the instance to register")
	sshRegisterCmd.Flags().StringVarP(&sshRegisterDto.User, "user", "u", "", "The user to use to login to the instance")
	sshRegisterCmd.Flags().IntVarP(&sshRegisterDto.Port, "port", "p", 22, "The port to connect to")
	sshCmd.AddCommand(sshRegisterCmd)
}
