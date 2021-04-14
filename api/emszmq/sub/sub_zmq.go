// //
// //  Subscriber
// //

// package sub

// import (
// 	"fmt"
// 	"context"
// 	"github.com/go-zeromq/zmq4"
// )

// //SubZmq subscriber
// func SubZmq() {
// 	//  Prepare our subscriber
// 	subscriber, _ := zmq4.NewSub(SUB)
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

// Copyright 2018 The go-zeromq Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// PubSub envelope subscriber
package emszmq

import (
	"context"
	"log"

	"github.com/go-zeromq/zmq4"
)

func main() {
	log.SetPrefix("psenvsub: ")

	//  Prepare our subscriber
	sub := zmq4.NewSub(context.Background())
	defer sub.Close()

	err := sub.Dial("tcp://localhost:5563")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}

	err = sub.SetOption(zmq4.OptionSubscribe, "B")
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
