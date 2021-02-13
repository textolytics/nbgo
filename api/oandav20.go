package main

import (
	"flag"
	"fmt"

	oanda "github.com/santegoeds/oanda"
)

var (
	tokenv20   = flag.String("token", "63d484ee1bf37009848ccdb6c421ad75-bd7b55136f1ffc1c486b6bafe33d87f6", "Oanda authorization token.")
	accountv20 = flag.Int64("account", 3914094, "Oanda account.")
)

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

	client.SelectAccount(oanda.Id(*account))

	// List available instruments
	instruments, err := client.Instruments(nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(instruments)

	// Buy one unit of EUR/USD with a trailing stop of 10 pips.
	tradeInfo, err := client.NewTrade(oanda.Buy, 1, "eur_usd", oanda.TrailingStop(10.0))
	if err != nil {
		panic(err)
	}
	fmt.Println(tradeInfo)

	// Create and run a price server.
	priceServer, err := client.NewPriceServer("eur_usd")
	if err != nil {
		panic(err)
	}
	priceServer.ConnectAndHandle(func(instrument string, tick oanda.PriceTick) {
		fmt.Println("Received tick:", instrument, tick)
		priceServer.Stop()
	})

	// Close the previously opened trade.
	tradeCloseInfo, err := client.CloseTrade(tradeInfo.TradeId)
	if err != nil {
		panic(err)
	}
	fmt.Println(tradeCloseInfo)
}
