package emszmqpb

import (
	"fmt"

	emszmqpb "github.com/pebbe/zmq4"
	"github.com/textolytics/nbgo/nbdw/dwclickhouse"
)

type ZmqSubKrakenTick string
type ZmqSubKrakenDepth string
type ZmqSubKrakenEURUSDTick string
type ZmqSubKrakenEURUSDDepth string

//SubKrakenTick subscriber
func SubKrakenTick() (ZmqSubKrakenTick string) {
	//  Prepare our subscriber
	subscriberSubKrakenTick, _ := emszmqpb.NewSocket(emszmqpb.SUB)
	// defer subscriber.Close()
	subscriberSubKrakenTick.Connect("tcp://zmq.nb.lan:5558")
	subscriberSubKrakenTick.SetSubscribe("")
	for {
		//  Read envelope with address
		// address, _ := subscriber.Recv(0)
		//  Read message contents
		ZmqSubKrakenTick, _ := subscriberSubKrakenTick.Recv(0)
		fmt.Println(ZmqSubKrakenTick)
		// spew.Dump("%+v\n", contents)
		// return address, contents
		// return ZmqSubKrakenTick
	}

}

//SubKrakenDepth subscriber
func SubKrakenDepth() (ZmqSubKrakenDepth string) {
	//  Prepare our subscriber
	subscriberSubKrakenDepth, _ := emszmqpb.NewSocket(emszmqpb.SUB)
	// fmt.Println(subscriber)
	// fmt.Println(err)
	// defer subscriberSubKrakenDepth.Close()
	subscriberSubKrakenDepth.Connect("tcp://192.168.0.13:5560")
	// if subscriber_err != nil {
	// 	fmt.Println(subscriber_err)
	// }
	subscriberSubKrakenDepth.SetSubscribe("")
	for {
		//  Read envelope with address
		// address, _ := subscriber.Recv(0)
		//  Read message contents
		ZmqSubKrakenDepth, _ := subscriberSubKrakenDepth.Recv(0)
		fmt.Println(ZmqSubKrakenDepth)
		// spew.Dump("%+v\n", contents)
		// return address, contents
		// return ZmqSubKrakenDepth
	}
}

//SubKrakenEURUSDTick subscriber
func SubKrakenEURUSDTick() ZmqSubKrakenEURUSDTick {
	//  Prepare our subscriber
	subscriberSubKrakenEURUSDTick, _ := emszmqpb.NewSocket(emszmqpb.SUB)
	// defer subscriber.Close()
	subscriberSubKrakenEURUSDTick.Connect("tcp://zmq.nb.lan:5559")
	subscriberSubKrakenEURUSDTick.SetSubscribe("")
	for {
		//  Read envelope with address
		// address, _ := subscriber.Recv(0)
		//  Read message contents
		ZmqSubKrakenEURUSDTick, _ := subscriberSubKrakenEURUSDTick.Recv(0)
		fmt.Println(ZmqSubKrakenEURUSDTick)
		// spew.Dump("%+v\n", contents)
		// return address, contents
	}
}

//SubKrakenEURUSDDepth subscriber
func SubKrakenEURUSDDepth() (ZmqSubKrakenEURUSDDepth string) {
	//  Prepare our subscriber
	subscriberSubKrakenEURUSDDepth, _ := emszmqpb.NewSocket(emszmqpb.SUB)
	// defer subscriber.Close()
	subscriberSubKrakenEURUSDDepth.Connect("tcp://zmq.nb.lan:5561")
	subscriberSubKrakenEURUSDDepth.SetSubscribe("")
	for {
		//  Read envelope with address
		// address, _ := subscriber.Recv(0)
		//  Read message contents
		ZmqSubKrakenEURUSDDepth, _ = subscriberSubKrakenEURUSDDepth.Recv(0)
		// fmt.Println(ZmqSubKrakenEURUSDDepth)
		// spew.Dump("%+v\n", contents)
		// return address, contents
	}
}

//
//  Reading from multiple sockets.
//  This version uses zmq.Poll()
//

func MultipleSubPoller() {
	//  Connect to task ventilator
	// receiver, _ := zmq.NewSocket(zmq.PULL)
	// defer receiver.Close()
	// receiver.Connect("tcp://localhost:5557")

	// //  Connect to weather server
	// subscriber, _ := zmq.NewSocket(zmq.SUB)
	// defer subscriber.Close()
	// subscriber.Connect("tcp://localhost:5556")
	// subscriber.SetSubscribe("")

	//  Connect to subscriberSubKrakenEURUSDDepth
	subscriberSubKrakenEURUSDDepth, _ := emszmqpb.NewSocket(emszmqpb.SUB)
	defer subscriberSubKrakenEURUSDDepth.Close()
	subscriberSubKrakenEURUSDDepth.Connect("tcp://zmq.nb.lan:5561")
	subscriberSubKrakenEURUSDDepth.SetSubscribe("")

	//subscriberSubKrakenEURUSDTick subscriber
	subscriberSubKrakenEURUSDTick, _ := emszmqpb.NewSocket(emszmqpb.SUB)
	defer subscriberSubKrakenEURUSDTick.Close()
	subscriberSubKrakenEURUSDTick.Connect("tcp://zmq.nb.lan:5559")
	subscriberSubKrakenEURUSDTick.SetSubscribe("")

	//subscriberSubKrakenDepth subscriber
	subscriberSubKrakenDepth, _ := emszmqpb.NewSocket(emszmqpb.SUB)
	defer subscriberSubKrakenDepth.Close()
	subscriberSubKrakenDepth.Connect("tcp://192.168.0.13:5560")
	subscriberSubKrakenDepth.SetSubscribe("")

	//subscriberSubKrakenTick subscriber
	subscriberSubKrakenTick, _ := emszmqpb.NewSocket(emszmqpb.SUB)
	defer subscriberSubKrakenTick.Close()
	subscriberSubKrakenTick.Connect("tcp://zmq.nb.lan:5558")
	subscriberSubKrakenTick.SetSubscribe("")

	//  Initialize DW Client
	dwclickhouseClient := dwclickhouse.ClickHouseDWClient()

	//  Initialize poll set
	poller := emszmqpb.NewPoller()
	// poller.Add(receiver, zmq.POLLIN)
	poller.Add(subscriberSubKrakenTick, emszmqpb.POLLIN)
	poller.Add(subscriberSubKrakenDepth, emszmqpb.POLLIN)
	poller.Add(subscriberSubKrakenEURUSDTick, emszmqpb.POLLIN)
	poller.Add(subscriberSubKrakenEURUSDDepth, emszmqpb.POLLIN)

	//  Process messages from both sockets
	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			// case receiver:
			// 	task, _ := s.Recv(0)
			// 	//  Process task
			// 	fmt.Println("Got task:", task)
			case subscriberSubKrakenTick:
				update, _ := s.Recv(0)
				//  Process weather update
				dwclickhouse.ClickHouseDWClientInsert(dwclickhouseClient, string(`subKrakenTick`), string(update))
				fmt.Println("Got SubKrakenTick update: ", update)

			case subscriberSubKrakenDepth:
				update, _ := s.Recv(0)
				//  Process weather update
				dwclickhouse.ClickHouseDWClientInsert(dwclickhouseClient, string(`subKrakenDepth`), string(update))
				fmt.Println("Got SubKrakenDepth update: ", update)

			case subscriberSubKrakenEURUSDTick:
				update, _ := s.Recv(0)
				//  Process weather update
				dwclickhouse.ClickHouseDWClientInsert(dwclickhouseClient, string(`subKrakenEURUSDTick`), string(update))
				fmt.Println("Got SubKrakenEURUSDTick update: ", update)

			case subscriberSubKrakenEURUSDDepth:
				update, _ := s.Recv(0)
				//  Process weather update
				dwclickhouse.ClickHouseDWClientInsert(dwclickhouseClient, "subKrakenEURUSDDepth", string(update))
				fmt.Println("Got SubKrakenEURUSDDepth update: ", update)

			}
		}
	}
}
