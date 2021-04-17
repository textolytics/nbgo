package emszmqpb

import (
	"fmt"

	emszmqpb "github.com/pebbe/zmq4"
)

//SubZmqTick subscriber
func SubZmqTick() {
	//  Prepare our subscriber
	subscriber, _ := emszmqpb.NewSocket(emszmqpb.Type(4))
	// defer subscriber.Close()
	subscriber.Connect("tcp://zmq.nb.lan:5558")
	subscriber.SetSubscribe("")
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
		// spew.Dump("%+v\n", contents)
		// return address, contents

	}
}

//SubZmqDepth subscriber
func SubZmqDepth() {
	//  Prepare our subscriber
	subscriber, _ := emszmqpb.NewSocket(emszmqpb.Type(4))
	// defer subscriber.Close()
	subscriber.Connect("tcp://zmq.nb.lan:5560")
	subscriber.SetSubscribe("")
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
		// spew.Dump("%+v\n", contents)
		// return address, contents

	}
}

//SubZmqEURUSDTick subscriber
func SubZmqEURUSDTick() {
	//  Prepare our subscriber
	subscriber, _ := emszmqpb.NewSocket(emszmqpb.Type(4))
	// defer subscriber.Close()
	subscriber.Connect("tcp://zmq.nb.lan:5559")
	subscriber.SetSubscribe("")
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
		// spew.Dump("%+v\n", contents)
		// return address, contents

	}
}

//SubZmqEURUSDDepth subscriber
func SubZmqEURUSDDepth() {
	//  Prepare our subscriber
	subscriber, _ := emszmqpb.NewSocket(emszmqpb.Type(4))
	// defer subscriber.Close()
	subscriber.Connect("tcp://zmq.nb.lan:5561")
	subscriber.SetSubscribe("")
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
		// spew.Dump("%+v\n", contents)
		// return address, contents

	}
}
