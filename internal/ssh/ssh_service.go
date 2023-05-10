package ssh

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var (
	sshService     SSHService
	sshServiceOnce sync.Once
)

type SSHService interface{}

func NewSshService(client *ec2.Client) SSHService {
	if sshService == nil {
		sshServiceOnce.Do(func() {
			sshService = &sshServiceImpl{}
		})
	}
	return sshService
}
