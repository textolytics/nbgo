// Copyright 2018 The go-zeromq Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// PubSub envelope subscriber
package emszmq

import (
	"context"
	"log"

	"github.com/go-zeromq/zmq4"
)

func SubZmqTick() {
	log.SetPrefix("psenvsub: ")

	//  Prepare our subscriber
	sub := zmq4.NewSub(context.Background())
	defer sub.Close()

	err := sub.Dial("tcp://zmq.nb.lan:5558")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}

	err = sub.SetOption(zmq4.OptionSubscribe, "")
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}

	for {
		// Read envelope
		msg, err := sub.Recv()
		if err != nil {
			log.Fatalf("could not receive message: %v", err)
		}
		log.Printf("[%s] %s\n", msg.Frames[0], msg.Frames[1])
	}
}

// import (
// 	"fmt"

// 	zmq "github.com/go-zeromq/zmq4"
// )

// //SubZmqTick subscriber
// func SubZmqTick() {
// 	//  Prepare our subscriber
// 	subscriber, _ := zmq.NewSocket(zmq.SUB)
// 	// defer subscriber.Close()
// 	subscriber.Connect("tcp://zmq.nb.lan:5558")
// 	subscriber.SetSubscribe("")
// 	for {
// 		//  Read envelope with address
// 		address, _ := subscriber.Recv(0)
// 		//  Read message contents
// 		contents, _ := subscriber.Recv(0)
// 		fmt.Printf("[%s] %s\n", address, contents)
// 		// spew.Dump("%+v\n", contents)
// 		// return address, contents

// 	}
// }

// //SubZmqDepth subscriber
// func SubZmqDepth() {
// 	//  Prepare our subscriber
// 	subscriber, _ := zmq.NewSocket(zmq.SUB)
// 	// defer subscriber.Close()
// 	subscriber.Connect("tcp://zmq.nb.lan:5560")
// 	subscriber.SetSubscribe("")
// 	for {
// 		//  Read envelope with address
// 		address, _ := subscriber.Recv(0)
// 		//  Read message contents
// 		contents, _ := subscriber.Recv(0)
// 		fmt.Printf("[%s] %s\n", address, contents)
// 		// spew.Dump("%+v\n", contents)
// 		// return address, contents

// 	}
// }

// //SubZmqEURUSDTick subscriber
// func SubZmqEURUSDTick() {
// 	//  Prepare our subscriber
// 	subscriber, _ := zmq.NewSocket(zmq.SUB)
// 	// defer subscriber.Close()
// 	subscriber.Connect("tcp://zmq.nb.lan:5559")
// 	subscriber.SetSubscribe("")
// 	for {
// 		//  Read envelope with address
// 		address, _ := subscriber.Recv(0)
// 		//  Read message contents
// 		contents, _ := subscriber.Recv(0)
// 		fmt.Printf("[%s] %s\n", address, contents)
// 		// spew.Dump("%+v\n", contents)
// 		// return address, contents

// 	}
// }

// //SubZmqEURUSDDepth subscriber
// func SubZmqEURUSDDepth() {
// 	//  Prepare our subscriber
// 	subscriber, _ := zmq.NewSocket(zmq.SUB)
// 	// defer subscriber.Close()
// 	subscriber.Connect("tcp://zmq.nb.lan:5561")
// 	subscriber.SetSubscribe("")
// 	for {
// 		//  Read envelope with address
// 		address, _ := subscriber.Recv(0)
// 		//  Read message contents
// 		contents, _ := subscriber.Recv(0)
// 		fmt.Printf("[%s] %s\n", address, contents)
// 		// spew.Dump("%+v\n", contents)
// 		// return address, contents

// 	}
// }
