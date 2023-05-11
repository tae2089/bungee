package ssh

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type sshServiceImpl struct {
	db *sql.DB
}

var _ SSHService = (*sshServiceImpl)(nil)

// RegisterSSHInfo implements SSHService
func (ssh *sshServiceImpl) RegisterSSHInfo(s SshRegisterDto) error {
	// insert data
	_, err := ssh.db.Exec("INSERT INTO ssh_info (name, user, host, port) VALUES (?,?,?,?)", s.Key, s.User, s.Host, s.Port)
	if err != nil {
		return err
	}
	return nil
}
