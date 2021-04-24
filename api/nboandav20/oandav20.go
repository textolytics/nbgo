package oandav20

// import (
// 	"flag"
// 	"fmt"

// 	oanda "github.com/santegoeds/oanda"
// )

// var (
// 	token   = flag.String("token", "63d484ee1bf37009848ccdb6c421ad75-bd7b55136f1ffc1c486b6bafe33d87f6", "Oanda authorization token.")
// 	account = flag.Int64("account", 3914094, "Oanda account.")
// )

// // Stream ,
// func Stream() {
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

// 	client.SelectAccount(oanda.Id(*account))

// 	// List available instruments
// 	instruments, err := client.Instruments(nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(instruments)

// 	// Buy one unit of EUR/USD with a trailing stop of 10 pips.
// 	tradeInfo, err := client.NewTrade(oanda.Buy, 1, "eur_usd", oanda.TrailingStop(10.0))
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(tradeInfo)

// 	// Create and run a price server.
// 	priceServer, err := client.NewPriceServer("eur_usd")
// 	if err != nil {
// 		panic(err)
// 	}
// 	priceServer.ConnectAndHandle(func(instrument string, tick oanda.PriceTick) {
// 		fmt.Println("Received tick:", instrument, tick)
// 		priceServer.Stop()
// 	})

// 	// Close the previously opened trade.
// 	tradeCloseInfo, err := client.CloseTrade(tradeInfo.TradeId)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(tradeCloseInfo)
// }

// package oandaV20

// import (
// 	"flag"
// 	"fmt"

// 	oanda "github.com/santegoeds/oanda"
// )

// var (
// 	token   = flag.String("token", "63d484ee1bf37009848ccdb6c421ad75-bd7b55136f1ffc1c486b6bafe33d87f6", "Oanda authorization token.")
// 	account = flag.Int64("account", 3914094, "Oanda account.")
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
