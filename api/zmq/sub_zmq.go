//
//  Subscriber
//

package zmq

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

//SubZmq subscriber
func SubZmq() (address string, contents string) {
	//  Prepare our subscriber
	subscriber, _ := zmq.NewSocket(zmq.Type(6))
	subscriber.SetSubscribe("1")
	// subscriber.GetEvents()
	// defer subscriber.Close()
	subscriber.Connect("tcp://192.168.0.13:5558")
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
