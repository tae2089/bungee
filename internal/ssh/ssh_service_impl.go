package ssh

type sshServiceImpl struct{}

var _ SSHService = (*sshServiceImpl)(nil)
