package main

import (
	"log"
	// Based on :
	socks5 "github.com/armon/go-socks5"
)

const (
	login    = "CHANGE_ME"
	password = "CHANGE_ME"
	ipport   = "0.0.0.0:1234"
	timeout  = 10
)

func main() {
	log.Printf("Proxy initializing...")
	// Make credencials
	creds := socks5.StaticCredentials{
		login: password,
	}
	log.Printf("Credentials passed...")
	c := socks5.UserPassAuthenticator{Credentials: creds}
	method := []socks5.Authenticator{c}
	// Make config
	conf := &socks5.Config{
		AuthMethods: method,
	}
	log.Printf("Config generated...")

	// Create instance
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	log.Printf("Proxy object created...Listening on %v", ipport)
	// Create SOCKS5 proxy

	if err := server.ListenAndServe("tcp", ipport); err != nil {
		panic(err)
	}

}
