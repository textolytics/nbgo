//
//  Subscriber
//

package sub

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

//SubZmq subscriber
func SubZmq() {
	//  Prepare our subscriber
	subscriber, _ := zmq.NewSocket(zmq.SUB)
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
