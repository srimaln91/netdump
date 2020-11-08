package ssh

import (
	"io"

	"golang.org/x/crypto/ssh"
)

type Session struct {
	*ssh.Session
}

func (session *Session) GetInterfaces() (stdout, stderr io.Reader, err error) {
	stdout, err = session.StdoutPipe()
	if err != nil {
		return
	}

	stderr, err = session.StderrPipe()
	if err != nil {
		return
	}

	return
}
