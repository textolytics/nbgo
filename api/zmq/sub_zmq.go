//
//  Subscriber
//

package zmq

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

//SubZmq subscriber
func SubZmq() {
	//  Prepare our subscriber
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5563")
	subscriber.SetSubscribe("B")

	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
	}
}
