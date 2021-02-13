package api

import (
	"log"

	"github.com/zeromq/goczmq"
)

func main() {
	// Create a router channeler and bind it to port 5555.
	// A channeler provides a thread safe channel interface
	// to a *Sock
	router := goczmq.NewRouterChanneler("tcp://*:5555")
	defer router.Destroy()

	log.Println("router created and bound")

	// Create a dealer channeler and connect it to the router.
	dealer := goczmq.NewDealerChanneler("tcp://127.0.0.1:5555")
	defer dealer.Destroy()

	log.Println("dealer created and connected")

	// Send a 'Hello' message from the dealer to the router.
	dealer.SendChan <- [][]byte{[]byte("Hello")}
	log.Println("dealer sent 'Hello'")

	// Receve the message as a [][]byte. Since this is
	// a router, the first frame of the message wil
	// be the routing frame.
	request := <-router.RecvChan
	log.Printf("router received '%s' from '%v'", request[1], request[0])

	// Send a reply. First we send the routing frame, which
	// lets the dealer know which client to send the message.
	router.SendChan <- [][]byte{request[0], []byte("World")}
	log.Printf("router sent 'World'")

	// Receive the reply.
	reply := <-dealer.RecvChan
	log.Printf("dealer received '%s'", string(reply[0]))
}

// <<<<<<< HEAD

// import (
// 	"flag"
// 	"fmt"

// 	oanda "github.com/santegoeds/oanda"
// )

// var (
// 	token   = flag.String("token", "63d484ee1bf37009848ccdb6c421ad75-bd7b55136f1ffc1c486b6bafe33d87f6", "Oanda authorization token.")
// 	account = flag.Int64("account", 101-004-3748257-001, "Oanda account.")
// 	instrs  []string
// )

// //OandaTicker is Oanda Price Tick
// type OandaTicker struct {
// 	TickerName string
// 	Tick       oanda.PriceTick
// }

// //GetOandaTickStream for quotes stream from Oanda route
// func (t OandaTicker) GetOandaTickStream(tickerName string, tick oanda.PriceTick) {
// 	t.TickerName = tickerName
// 	t.Tick = tick
// 	// return OandaTicker
// }

// //OandaTickStream get all quotes from route
// func OandaTickStream() {
// 	flag.Parse()
// 	if *token == "" {
// 		panic("An Oanda authorization token is required")
// 	}

// 	if *account == 0 {
// 		panic("An Oanda account is required")
// 	}

// 	client, err := oanda.NewFxPracticeClient(*token)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// List available account
// 	client.SelectAccount(oanda.Id(*account))

// 	// List available instruments
// 	instruments, err := client.Instruments(nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i := range instruments {
// 		fmt.Println(i)
// 		instrs = append(instrs, i)
// 	}

// 	// Create and run a NewPriceServer server.
// 	priceServer, err := client.NewPriceServer(instrs...)
// 	if err != nil {
// 		panic(err)
// 	}

// 	priceServer.ConnectAndHandle(func(instrs string, tick oanda.PriceTick) {
// 		if err != nil {
// 			fmt.Println("Received err:", err)
// 			panic(err)
// 		}
// 		go fmt.Println("Received tick:", instrs, tick)
// 		fmt.Printf("Received instrs type: %T | Tick type: %T \r\n ", instrs, tick)

// 		// func (TickerName instrs) () {

// 		// }
// 		OandaTicker.GetOandaTickStream(OandaTicker{}, instrs, tick)
// 		// fmt.Println(TickerTime, TickerBid, TickerAsk, TickerStatus)

// 		// }
// 		//		writeParquetTEST(&tickParquet)

// 		// priceServer.Stop()
// 		// return TickerName, Tick

// 	})
// 	// return TickerName, Tick
// }
// =======
// >>>>>>> b4a9927e20e2400646effbf2980a1851d8795b54
