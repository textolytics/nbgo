package main

import (
	"encoding/json"
	"log"
	"time"

	gateiows "github.com/gateio/gatews/go"
)

func ConnectGateioWs() {
	// create WsService with ConnConf, this is recommended, key and secret will be needed by some channels
	// ctx and logger could be nil, they'll be initialized by default
	Ws, err := gateiows.NewWsService(nil, nil, gateiows.NewConnConfFromOption(&gateiows.ConfOptions{
		Key:           "38e92d81c9612ed8c46bbae3484c06a4",
		Secret:        "44f78afbb7054ceffd8a87eef39e01cf39551315a8245531447b299c46ccaeaa",
		MaxRetryConn:  10, // default value is math.MaxInt64, set it when needs
		SkipTlsVerify: false,
	}))
	// we can also do nothing to get a WsService, all parameters will be initialized by default and default url is spot
	// but some channels need key and secret for auth, we can also use set function to set key and secret
	// ws, err := gate.NewWsService(nil, nil, nil)
	// ws.SetKey("YOUR_API_KEY")
	// ws.SetSecret("YOUR_API_SECRET")
	if err != nil {
		log.Printf("NewWsService err:%s", err.Error())
		return
	}

	// checkout connection status when needs
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			log.Println("connetion status:", Ws.Status())
		}
	}()

}

func GetLocalDepthSnapshot(instruments []string, Ws *gateiows.WsService) error {

	// create callback functions for receive messages
	callOrder := gateiows.NewCallBack(func(msg *gateiows.UpdateMsg) {
		// parse the message to struct we need
		var order []gateiows.SpotOrderMsg
		if err := json.Unmarshal(msg.Result, &order); err != nil {
			log.Printf("order Unmarshal err:%s", err.Error())
		}
		log.Printf("%+v", order)
	})

	callTrade := gateiows.NewCallBack(func(msg *gateiows.UpdateMsg) {
		var trade gateiows.SpotTradeMsg
		if err := json.Unmarshal(msg.Result, &trade); err != nil {
			log.Printf("trade Unmarshal err:%s", err.Error())
		}
		log.Printf("%+v", trade)
	})

	// first, we need set callback function
	Ws.SetCallBack(gateiows.ChannelSpotOrder, callOrder)
	Ws.SetCallBack(gateiows.ChannelSpotPublicTrade, callTrade)
	// second, after set callback function, subscribe to any channel you are interested into
	if err := Ws.Subscribe(gateiows.ChannelSpotPublicTrade, instruments); err != nil {
		log.Printf("Subscribe err:%s", err.Error())
		return err
	}
	if err := Ws.Subscribe(gateiows.ChannelSpotBookTicker, instruments); err != nil {
		log.Printf("Subscribe err:%s", err.Error())
		return err
	}
	// example for maintaining local order book
	// gateiows.LocalOrderBook(context.Background(), Ws, []string{"BTC_USDT"})

}
