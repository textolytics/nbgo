//
//  Subscriber
//

package zmq

import (
	"fmt"

	zmq4 "github.com/pebbe/zmq4"
)

//SubZmq subscriber
func SubZmq() (address string, contents string) {
	//  Prepare our subscriber
	subscriber, _ := zmq4.NewSocket(zmq4.SUB)
	subscriber.SetSubscribe("1")
	// subscriber.GetEvents()
	// defer subscriber.Close()
	subscriber.Connect("tcp://192.168.0.13:5560")
	fmt.Printf("[%s] %s\n", subscriber, subscriber)
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
		// spew.Dump("%+v\n", contents)
		return address, contents
	}
}
