//
//  Pubsub envelope subscriber.
//

package emszmq

import (
	"fmt"

	"github.com/go-zeromq/zmq4"
)

//SubKrakenTick subscriber
func SubKrakenTick() {
	//  Prepare our subscriber
	subscriber, _ := zmq4.NewSocket(zmq4.Type(zmq4.SUB))
	// defer subscriber.Close()
	subscriber.Connect("tcp://zmq.nb.lan:5560")
	subscriber.SetSubscribe("")
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
	}
}
