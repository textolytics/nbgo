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
	subscriber.SetSubscribe("B")
	defer subscriber.Close()
	subscriber.Connect("tcp://192.168.0.13:5563")
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
	}
}
