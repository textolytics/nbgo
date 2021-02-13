//
//  Publisher
//

package zmq

import (
	zmq4 "github.com/pebbe/zmq4"
)

// PubZmq Publisher
func PubZmq() {
	//  Prepare our publisher
	publisher, _ := zmq4.NewSocket(zmq4.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://*:5563")
	for {
		//  Write two messages, each with an envelope and content
		publisher.Send("A", zmq4.SNDMORE)
		publisher.Send("We don't want to see this", 0)
		publisher.Send("B", zmq4.SNDMORE)
		publisher.Send("We would like to see this", 0)
		// time.Sleep(time.Second)
	}

}
