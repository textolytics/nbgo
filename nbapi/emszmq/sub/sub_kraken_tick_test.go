//
//  Pubsub envelope subscriber.
//

package emszmq

// import (
// 	"testing"

// 	zmq "github.com/go-zeromq/zmq4"
// )

// //SubKrakenTick PAt Subs
// func TestSubKrakenTick(t *testing.T) {
// 	//  Prepare our subscriber
// 	subscriber, _ := zmq.NewSocket(zmq.Type(6))
// 	defer subscriber.Close()
// 	subscriber.Connect("tcp://zmq.nb.lan:5558")
// 	subscriber.SetSubscribe("")

// 	for i := 0; i < 1000; i++ {
// 		//  Read envelope with address
// 		address, _ := subscriber.Recv(0)
// 		//  Read message contents
// 		contents, _ := subscriber.Recv(0)
// 		// fmt.Printf("[%s] %s\n", address, contents)
// 		t.Error(address, contents)
// 	}

// }
