//
//  Publisher
//

package zmq

import (
	zmq "github.com/pebbe/zmq4"
)

// PubZmq Publisher
func PubZmq() {
	//  Prepare our publisher
	publisher, _ := zmq.NewSocket(zmq.Type(zmq.PUB))
	defer publisher.Close()
	publisher.Bind("tcp://*:5500")
	for {
		//  Write two messages, each with an envelope and content
		publisher.Send("A", zmq.Flag(zmq.SNDMORE))
		publisher.Send("We don't want to see this", 0)
		publisher.Send("B", zmq.Flag(zmq.SNDMORE))
		publisher.Send("We would like to see this", 0)
		// time.Sleep(time.Second)
	}

}
