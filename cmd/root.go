package cmd

import (
	"fmt"
	"os"

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
)

func init() {
	// cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
