package auth

import (
	"net"
	"os"

	sshcon "github.com/srimaln91/netdump/connection/ssh"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type HostAgent struct {
	config *ssh.ClientConfig
}

func NewSSHAgentProvider(user string) sshcon.AuthProvider {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			SSHAgent(),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	au := HostAgent{
		config: sshConfig,
	}

	return &au
}

func (agent *HostAgent) GetClientConfig() *ssh.ClientConfig{
	return agent.config
}


func SSHAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}