package ssh

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
)

type sshServiceImpl struct {
	db *sql.DB
}

var _ SSHService = (*sshServiceImpl)(nil)

// RegisterSSHInfo implements SSHService
func (ssh *sshServiceImpl) RegisterSSHInfo(s *SshRegisterDto, privateKey string) error {
	// insert data
	defer ssh.db.Close()
	_, err := ssh.db.Exec("INSERT INTO ssh_info (name, user, host, port, key, password) VALUES (?,?,?,?,?,?)", s.Name, s.User, s.Host, s.Port, privateKey, s.Password)
	if err != nil {
		return err
	}
	return nil
}

// ShowSSHInfoList implements SSHService
func (ssh *sshServiceImpl) ShowSSHInfoList() error {
	// select data
	defer ssh.db.Close()
	rows, err := ssh.db.Query("SELECT name, user, host, port, password FROM ssh_info")
	if err != nil {
		return err
	}
	defer rows.Close()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "User", "Host", "Port", "Password"})
	for rows.Next() {
		var s SshRegisterDto
		err := rows.Scan(&s.Name, &s.User, &s.Host, &s.Port, &s.Password)
		if err != nil {
			return err
		}
		table.Append([]string{
			s.Name, s.User, s.Host, fmt.Sprint(s.Port), s.Password,
		})
	}
	table.Render()
	return nil
}
