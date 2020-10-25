package ssh

import (
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type SSHConn struct {
	Address string
	// clientConfig *ssh.ClientConfig
	client  *ssh.Client
	session *ssh.Session
}

func (conn *SSHConn) Connect() (err error) {

	sshConfig := &ssh.ClientConfig{
		User: "srimal",
		Auth: []ssh.AuthMethod{
			SSHAgent(),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	c, err := ssh.Dial("tcp", "35.247.154.95:22", sshConfig)

	if err != nil {
		return
	}

	// Create session
	session, err := c.NewSession()
	if err != nil {
		return
	}

	conn.client = c
	conn.session = session

	return
}

func SSHAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}

func (conn *SSHConn) NewSession() (*Session, error) {

	sess, err := conn.client.NewSession()
	if err != nil {
		return nil, err
	}

	return &Session{
		sess,
	}, nil
}
