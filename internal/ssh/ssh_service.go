package ssh

import (
	"database/sql"
	"sync"
)

var (
	sshService     SSHService
	sshServiceOnce sync.Once
)

type SSHService interface {
	RegisterSSHInfo(s *SshRegisterDto, privateKey string) error
	ShowSSHInfoList() error
}

func NewSshService(db *sql.DB) SSHService {
	if sshService == nil {
		sshServiceOnce.Do(func() {
			sshService = &sshServiceImpl{
				db: db,
			}
		})
	}
	return sshService
}

type SshRegisterDto struct {
	Host     string
	Port     int
	User     string
	Name     string
	Key      string
	Password string
}
