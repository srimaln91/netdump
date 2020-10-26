package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/srimaln91/netdump/connection/ssh"
	"github.com/srimaln91/netdump/connection/ssh/auth"
)

func main() {

	auth := auth.NewSSHAgentProvider("srimal")

	con := &ssh.SSHConn{}
	err := con.Connect("35.247.154.95:22", auth)

	if err != nil {
		fmt.Println(err)
		return
	}

	session, err := con.NewSession()
	if err != nil {
		fmt.Println(err)
		return
	}

	stdout, _, err := session.GetInterfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	go func(){
		err := session.RunCommand(`/usr/bin/sudo tcpdump -i eth0 -s 0 -A 'tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x47455420 or tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x504F5354 or tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x48545450 or tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x3C21444F'`)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	go func(){
		io.Copy(os.Stdout, stdout)
	}()

	// go io.Copy(os.Stdout, stdout)
	fmt.Println("Starting writers")

	// handle route using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		go func(){
			_, err := io.Copy(w, stdout)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}()

		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		for {
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			} else {
				log.Println("Damn, no flush")
			}
			time.Sleep(2 * time.Second)
		}

	})

	// listen to port
	http.ListenAndServe("0.0.0.0:5050", nil)
}
