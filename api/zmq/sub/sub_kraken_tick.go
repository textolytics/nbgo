//
//  Pubsub envelope subscriber.
//

package sub

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

//SubKrakenTick subscriber
func SubKrakenTick() {
	//  Prepare our subscriber
	subscriber, _ := zmq.NewSocket(zmq.SUB)
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
