package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tae2089/bungee/internal/di"
	"github.com/tae2089/bungee/internal/ssh"
)

var (
	// rootCmd represents the base command when called without any sub-commands
	sshRegisterDto = &ssh.SshRegisterDto{}
	keyPath        string
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
			// 사용자의 홈 디렉토리 경로 가져오기
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("홈 디렉토리 경로를 가져오는 데 실패했습니다. %s", err.Error())
			}
			// Private Key 파일 경로 생성
			keyPath = filepath.Join(homeDir, ".ssh/id_rsa")
			sshRegisterDto.Key = keyPath
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			sshService := di.InitSshService()
			if sshService == nil {
				return fmt.Errorf("Failed to initialize ssh service")
			}

			keyBytes, err := ioutil.ReadFile(sshRegisterDto.Key)
			if err != nil {
				return err
			}
			key := bytes.NewBuffer(keyBytes).String()
			err = sshService.RegisterSSHInfo(sshRegisterDto, key)
			return err
		},
		SilenceUsage: true,
	}
)

func init() {
	sshRegisterCmd.Flags().StringVarP(&sshRegisterDto.Host, "host", "H", "", "The host to connect to")
	sshRegisterCmd.Flags().StringVarP(&sshRegisterDto.Name, "Name", "n", "", "The name of the instance to register")
	sshRegisterCmd.Flags().StringVarP(&sshRegisterDto.User, "user", "u", "", "The user to use to login to the instance")
	sshRegisterCmd.Flags().IntVarP(&sshRegisterDto.Port, "port", "p", 22, "The port to connect to")
	sshRegisterCmd.Flags().StringVarP(&sshRegisterDto.Password, "password", "P", "", "The password to use to login to the instance")
	sshRegisterCmd.Flags().StringVarP(&sshRegisterDto.Key, "key", "k", "~/.ssh/id_rsa", "The private key to use to login to the instance")
	sshCmd.AddCommand(sshRegisterCmd)
}
