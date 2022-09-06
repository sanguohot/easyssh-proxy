package main

import (
	"fmt"
	"time"

	"github.com/appleboy/easyssh-proxy"
)

func main() {
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		Server:   "10.7.69.2",
		User:     "root",
		Password: "Jry2a12345678E%",
		Port:     "33022",
		Timeout:  60 * time.Second,
		Proxy: easyssh.DefaultConfig{
			User:     "2924663qad8298990",
			Server:   "www.baidu.com",
			Port:     "22011",
			Password: "Arn8JabUUHHh2UZ2LL",
			Type:     "socks5",
		},
	}

	// Call Run method with command you want to run on remote server.
	stdoutChan, stderrChan, doneChan, errChan, err := ssh.Stream("for i in {1..5}; do echo ${i}; cat /etc/os-release; sleep 1; done; exit 2;", 60*time.Second)
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		// read from the output channel until the done signal is passed
		isTimeout := true
	loop:
		for {
			select {
			case isTimeout = <-doneChan:
				break loop
			case outline := <-stdoutChan:
				fmt.Println("out:", outline)
			case errline := <-stderrChan:
				fmt.Println("err:", errline)
			case err = <-errChan:
			}
		}

		// get exit code or command error.
		if err != nil {
			fmt.Println("err: " + err.Error())
		}

		// command time out
		if !isTimeout {
			fmt.Println("Error: command timeout")
		}
	}
}
