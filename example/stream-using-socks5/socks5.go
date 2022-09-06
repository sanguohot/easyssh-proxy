package main

import (
	"fmt"
	"log"

	"github.com/haochen233/socks5"
)

func main() {
	// create socks client
	clnt := socks5.Client{
		ProxyAddr: "www.baidu.com:22011",
		// Authenticator supported by the client.
		// It must not be nil.
		Auth: map[socks5.METHOD]socks5.Authenticator{
			// If client want send NO_AUTHENTICATION_REQUIRED method to server, must
			// add socks5.NoAuth authenticator explicitly
			socks5.NO_AUTHENTICATION_REQUIRED: &socks5.NoAuth{},
			socks5.USERNAME_PASSWORD:          &socks5.UserPasswd{Username: "2924663qad8298990", Password: "Arn8JUUdaaah2UZLL"},
		},
	}

	// client send CONNECT command and get a tcp connection.
	// and use this connection transit data between you and www.google.com:80.
	conn, err := clnt.Connect(socks5.Version5, "10.8.69.2:33022")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
	// close connection.
	conn.Close()
}
