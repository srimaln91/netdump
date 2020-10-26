package ssh

import (
	"golang.org/x/crypto/ssh"
)

type SSHConn struct {
	Address string
	// clientConfig *ssh.ClientConfig
	client  *ssh.Client
	session *ssh.Session
}

type AuthProvider interface {
	GetClientConfig() *ssh.ClientConfig
}

func (conn *SSHConn) Connect(address string, authProvider AuthProvider) (err error) {

	c, err := ssh.Dial("tcp", address, authProvider.GetClientConfig())

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


func (conn *SSHConn) NewSession() (*Session, error) {

	sess, err := conn.client.NewSession()
	if err != nil {
		return nil, err
	}

	return &Session{
		sess,
	}, nil
}
