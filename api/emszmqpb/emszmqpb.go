package emszmqpb

import (
	"fmt"

	emszmqpb "github.com/pebbe/zmq4"
)

//SubZmqTick subscriber
func SubZmqTick() {
	//  Prepare our subscriber
	subscriber, _ := emszmqpb.NewSocket(emszmqpb.SUB)
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
	subscriber, err := emszmqpb.NewSocket(emszmqpb.SUB)
	fmt.Println(subscriber)
	fmt.Println(err)
	defer subscriber.Close()
	subscriber_err := subscriber.Connect("tcp://192.168.0.13:5560")
	if subscriber_err != nil {
		fmt.Println(subscriber_err)
	}
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
	subscriber, _ := emszmqpb.NewSocket(emszmqpb.XSUB)
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
	subscriber, _ := emszmqpb.NewSocket(emszmqpb.XSUB)
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
