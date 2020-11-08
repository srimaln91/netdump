package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/cors"
	"github.com/srimaln91/netdump/connection/ssh"
	"github.com/srimaln91/netdump/connection/ssh/auth"
)

type httpError struct{
	Data []string `json:"data"`
	Error bool `json:"error"`
	Message string `json:"message"`
}

func main() {

	host := flag.String("host", "localhost", "IP:port")
	user := flag.String("user", "", "User name")
	flag.Parse()

	auth := auth.NewSSHAgentProvider(*user)

	con := &ssh.SSHConn{}
	err := con.Connect(*host, auth)

	if err != nil {
		fmt.Println(err)
		return
	}

	session, err := con.NewSession()
	if err != nil {
		fmt.Println(err)
		return
	}

	handler := http.NewServeMux()

	var stdout io.Reader

	// handle route using handler function
	handler.HandleFunc("/log_stream", func(w http.ResponseWriter, r *http.Request) {

		if stdout == nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			return
		}

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
			time.Sleep(1 * time.Second)
		}

	})

	handler.HandleFunc("/apply_config", func(w http.ResponseWriter, r *http.Request) {

		session.Close()
		session, err := con.NewSession()
		if err != nil {
			fmt.Println(err)
			return
		}

		type Config struct {
			NetInterface string `json:"interface"`
		}

		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		var c Config
		err = json.Unmarshal(body, &c)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(c)

		stdout, _, err = session.GetInterfaces()
		if err != nil {
			fmt.Println(err)
			return
		}

		command := fmt.Sprintf(`
			/usr/bin/sudo tcpdump -i %s -s 0 -A 'tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x47455420 or tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x504F5354 or tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x48545450 or tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x3C21444F'`,
			c.NetInterface,
		)
		go func(){
			err := session.Run(command)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}()

		// go func(){
		// 	io.Copy(os.Stdout, stdout)
		// }()

		fmt.Println("Starting writers")

	})

	handler.HandleFunc("/interfaces", func(w http.ResponseWriter, r *http.Request) {

		session, err := con.NewSession()
		if err != nil {
			fmt.Println(err)
			return
		}

		defer session.Close()

		out, err := session.Output(`sudo ip r show | grep "src" | cut -d " " -f 3`)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		if len(out) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		interfaces:= strings.Split(string(out), "\n")

		resp := new(httpError)
		resp.Data = interfaces
		resp.Error = false
		
		respJSON, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
		
		return

	})

	//Server static files
	fs := http.FileServer(http.Dir("./web/build/"))

	handler.Handle("/", fs)

	fmt.Println("Navigate to localhost:5050")

	// listen to port
	http.ListenAndServe("0.0.0.0:5050", cors.Default().Handler(handler))
}
