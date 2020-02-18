package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/go-routeros/routeros"
)

var (
	command  = flag.String("command", "/ip service print", "RouterOS command")
	address  = flag.String("address", "192.168.0.1:48291", "RouterOS address and port")
	username = flag.String("username", "sdreep", "User name")
	password = flag.String("password", "chupakabra", "Password")
	async    = flag.Bool("async", false, "Use async code")
	useTLS   = flag.Bool("tls", false, "Use TLS")
)

func dial() (*routeros.Client, error) {
	if *useTLS {
		return routeros.DialTLS(*address, *username, *password, nil)
	}
	return routeros.Dial(*address, *username, *password)
}

func main() {
	flag.Parse()

	c, err := dial()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	defer c.Close()

	if *async {
		c.Async()
	}

	r, err := c.RunArgs(strings.Split(*command, " "))
	if err != nil {
		log.Fatal(err)
	}

	log.Print(r)
}
