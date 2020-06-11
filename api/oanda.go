package api

import (
	"flag"
	"fmt"

	"github.com/santegoeds/oanda"
)

var (
	token   = flag.String("token", "5b2e1521432ad31ef69270b682394010-4df302be03bbefb18ad70e457f3db869", "Oanda authorization token.")
	account = flag.Int64("account", 3914094, "Oanda account.")
	instrs  []string
)

type NbOandaTick struct {
	instrs string
	tick   string
}

type I interface {
	M()
}

func main() {
	flag.Parse()
	if *token == "" {
		panic("An Oanda authorization token is required")
	}

	if *account == 0 {
		panic("An Oanda account is required")
	}

	client, err := oanda.NewFxPracticeClient(*token)
	if err != nil {
		panic(err)
	}
	// List available account

	client.SelectAccount(oanda.Id(*account))

	// List available instruments
	instruments, err := client.Instruments(nil, nil)
	if err != nil {
		panic(err)
	}
	// fmt.Println(instruments)

	for i := range instruments {
		fmt.Println(i)
		instrs = append(instrs, i)
	}

	// Create and run a NewPriceServer server.
	priceServer, err := client.NewPriceServer(instrs...)
	if err != nil {
		panic(err)
	}

	priceServer.ConnectAndHandle(func(instrs string, tick oanda.PriceTick) {
		if err != nil {
			fmt.Println("Received err:", err)
			panic(err)
		}
		fmt.Println("Received tick:", instrs, tick)
		fmt.Printf("Received instrs type : %T tick type : %T ", instrs, tick)
		// tickParquet := oanda.PriceTick{
		// 	Time:   tick.Time,
		// 	Bid:    tick.Bid,
		// 	Ask:    tick.Ask,
		// 	Status: tick.Status,
		// }
		//		writeParquetTEST(&tickParquet)

		// priceServer.Stop()
	})
}
