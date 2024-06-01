package gateio

// type RedirectHandler interface {
// 	Get(http.ResponseWriter, *http.Request)
// 	Post(http.ResponseWriter, *http.Request)
// }

// Connect Client
// Get Instruments
// Get Instrument Details

// Get Tick
// Get Depth
// Send Order

type GateIoWs interface {
	GetTick()
	GetDepth()
	SendOrder()
}

type GateiowsApi struct {
	// strategyService nbgo.Strategy
}

// package api

// import (
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/go-chi/chi"
// 	"github.com/pkg/errors"

// 	js "github.com/tensor-programming/hex-microservice/serializer/json"
// 	ms "github.com/tensor-programming/hex-microservice/serializer/msgpack"
// 	"github.com/tensor-programming/hex-microservice/shortener"
// )

// type RedirectHandler interface {
// 	Get(http.ResponseWriter, *http.Request)
// 	Post(http.ResponseWriter, *http.Request)
// }

// type handler struct {
// 	redirectService shortener.RedirectService
// }

// func NewHandler(redirectService shortener.RedirectService) RedirectHandler {
// 	return &handler{redirectService: redirectService}
// }

// func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
// 	w.Header().Set("Content-Type", contentType)
// 	w.WriteHeader(statusCode)
// 	_, err := w.Write(body)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func (h *handler) serializer(contentType string) shortener.RedirectSerializer {
// 	if contentType == "application/x-msgpack" {
// 		return &ms.Redirect{}
// 	}
// 	return &js.Redirect{}
// }

// func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
// 	code := chi.URLParam(r, "code")
// 	redirect, err := h.redirectService.Find(code)
// 	if err != nil {
// 		if errors.Cause(err) == shortener.ErrRedirectNotFound {
// 			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 			return
// 		}
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently)
// }

// func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
// 	contentType := r.Header.Get("Content-Type")
// 	requestBody, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	redirect, err := h.serializer(contentType).Decode(requestBody)
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	err = h.redirectService.Store(redirect)
// 	if err != nil {
// 		if errors.Cause(err) == shortener.ErrRedirectInvalid {
// 			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 			return
// 		}
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	responseBody, err := h.serializer(contentType).Encode(redirect)
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	setupResponse(w, contentType, responseBody, http.StatusCreated)
// }

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	gatews "github.com/gateio/gatews/go"
// )

// // var GATEIO_API_KEY, _ = config.String("exchanges.gateio_api_key")
// // var GATEIO_API_SECRET, _ = config.String("exchanges.gateio_api_secret")

// var GATEIO_API_KEY = "38e92d81c9612ed8c46bbae3484c06a4"
// var GATEIO_API_SECRET = "44f78afbb7054ceffd8a87eef39e01cf39551315a8245531447b299c46ccaeaa"

// type GateIoEvent struct {
// }

// // func GateIoConfigInit(s string, b bool) (c gatews.ConnConf, err error) {
// // 	if err == nil {
// // 		return
// // 	}

// // 	return c, err
// // }

// var GateIoWSConnService, _ = gatews.NewWsService(context.TODO(), nil, gatews.NewConnConfFromOption(&gatews.ConfOptions{
// 	Key:           GATEIO_API_KEY,
// 	Secret:        GATEIO_API_SECRET,
// 	MaxRetryConn:  10, // default value is math.MaxInt64, set it when needs
// 	SkipTlsVerify: false,
// }))

// // GateWSConnService, err := gatews.NewWsService(context.TODO(), nil, gatews.NewConnConfFromOption(&gatews.ConfOptions

// // var GateIoWSConnService, _ = gatews.NewWsService(context.TODO(), nil, GateIoConnConf)

// type GateIoMarketData struct {
// 	GateIoWSConnService *gatews.WsService
// }

// type GateIoMarketDataSession interface {
// 	// GetChannels(c GetChannels) []string
// 	// GetChannelMarkets(m GetChannelMarkets) []string
// 	Level1(t gatews.SpotTradeMsg) *gatews.SpotTradeMsg
// 	Level2(b *gatews.SpotUpdateAllDepthMsg) *gatews.SpotUpdateAllDepthMsg
// 	Level2Update(bu *gatews.SpotUpdateDepthMsg) *gatews.SpotUpdateDepthMsg
// 	Level3(o []gatews.SpotOrderMsg) *[]gatews.SpotOrderMsg
// }

// type GateIoExecutionSession interface {
// }

// type GateIoExposureSession interface {
// }

// type GateIoTransferSession interface {
// }

// func GateIoMarketDataService() {

// 	// func getTrade() {
// 	callTrade := gatews.NewCallBack(func(msg *gatews.UpdateMsg) {
// 		var trade gatews.SpotTradeMsg
// 		if err := json.Unmarshal(msg.Result, &trade); err != nil {
// 			log.Printf("trade Unmarshal err:%s", err.Error())
// 		}
// 		// log.Println(trade)
// 		log.Printf("TRADE %+v", trade)
// 	})

// 	GateIoWSConnService.SetCallBack(gatews.ChannelSpotPublicTrade, callTrade)

// 	// second, after set callback function, subscribe to ChannelSpotPublicTrade any channel you are interested into
// 	if err := GateIoWSConnService.Subscribe(gatews.ChannelSpotPublicTrade, []string{"BTC_USDT", "ETH_USDT", "XRP_USDT", "LTC_BTC", "XRP_BTC"}); err != nil {
// 		log.Printf("Subscribe err:%s", err.Error())
// 		return
// 	}
// 	// }

// 	// func getOrder() {

// 	// create callback functions for receive messages
// 	callOrder := gatews.NewCallBack(func(msg *gatews.UpdateMsg) {
// 		// parse the message to struct we need
// 		var order gatews.OrderMsg
// 		if err := json.Unmarshal(msg.Result, &order); err != nil {
// 			log.Printf("order Unmarshal err:%s", err.Error())
// 		}
// 		log.Println(order)
// 		log.Printf("ORDER %+v", order)
// 	})

// 	GateIoWSConnService.SetCallBack(gatews.ChannelSpotOrder, callOrder)

// 	// second, after set callback function, subscribe to ChannelSpotBookTicker any channel you are interested into
// 	if err := GateIoWSConnService.Subscribe(gatews.ChannelSpotOrder, []string{"BTC_USDT", "ETH_USDT", "XRP_USDT", "LTC_BTC", "XRP_BTC"}); err != nil {
// 		log.Printf("Subscribe err:%s", err.Error())
// 		return
// 	}
// 	// return
// 	// }

// 	// func getOrderBook() (orderBook gatews.SpotUpdateAllDepthMsg) {
// 	callOrderBook := gatews.NewCallBack(func(msg *gatews.UpdateMsg) {
// 		var depth gatews.SpotUpdateAllDepthMsg
// 		if err := json.Unmarshal(msg.Result, &depth); err != nil {
// 			log.Printf("depth Unmarshal err:%s", err.Error())
// 		}
// 		log.Println(depth)
// 		log.Printf("DEPTH %+v", depth)
// 	})

// 	GateIoWSConnService.SetCallBack(gatews.ChannelSpotOrderBook, callOrderBook)

// 	// second, after set callback function, subscribe to ChannelSpotOrderBook any channel you are interested into
// 	if err := GateIoWSConnService.Subscribe(gatews.ChannelSpotOrderBook, []string{"BTC_USDT", "ETH_USDT", "XRP_USDT", "LTC_BTC", "XRP_BTC"}); err != nil {
// 		log.Printf("Subscribe err:%s", err.Error())
// 		return
// 	}
// 	// return orderBook
// 	// }

// 	// func getOrderBookUpdate() {
// 	callOrderBookUpdate := gatews.NewCallBack(func(msg *gatews.UpdateMsg) {
// 		var depthUpdate gatews.SpotUpdateDepthMsg
// 		if err := json.Unmarshal(msg.Result, &depthUpdate); err != nil {
// 			log.Printf("depthUpdate Unmarshal err:%s", err.Error())
// 		}
// 		log.Println(depthUpdate)
// 		log.Printf("DEPTH-UPDATE %+v", depthUpdate)
// 	})

// 	GateIoWSConnService.SetCallBack(gatews.ChannelSpotOrderBookUpdate, callOrderBookUpdate)

// 	// second, after set callback function, subscribe to ChannelSpotOrderBookUpdate any channel you are interested into
// 	if err := GateIoWSConnService.Subscribe(gatews.ChannelSpotOrderBookUpdate, []string{"BTC_USDT", "ETH_USDT", "XRP_USDT", "LTC_BTC", "XRP_BTC"}); err != nil {
// 		log.Printf("Subscribe err:%s", err.Error())
// 		return
// 	}

// 	// }

// 	ch := make(chan os.Signal, 8)
// 	signal.Ignore(syscall.SIGPIPE, syscall.SIGALRM)
// 	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT)
// 	// signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT, syscall.SIGKILL)
// 	<-ch
// }

// // func main() {

// // 	GateIoMarketDataService()
// // 	// log.Printf(GateIoMarketDataSession{item})

// // 	// clickhouse.ClickHouseConnect()
// // 	// gatews.GateIoLastSale()
// // 	// gatews.GateIoMarketDataSession.Level2()
// // } // ch := make(chan os.Signal, 8)
// // // signal.Ignore(syscall.SIGPIPE, syscall.SIGALRM)
// // // signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT)
// // // // signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT, syscall.SIGKILL)
// // <-ch

// // func GateIoLastSale() {

// // create WsService with ConnConf, this is recommended, key and secret will be needed by some channels
// // ctx and logger could be nil, they'll be initialized by default
// // GateWSConnService, err := gatews.NewWsService(context.TODO(), nil, gatews.NewConnConfFromOption(&gatews.ConfOptions{
// // 	Key:           GATEIO_API_KEY,
// // 	Secret:        GATEIO_API_SECRET,
// // 	MaxRetryConn:  10, // default value is math.MaxInt64, set it when needs
// // 	SkipTlsVerify: false,
// // }))
// // if err != nil {
// // 	log.Printf("NewWsService err:%s", err.Error())
// // 	return
// // }

// // callTrade := gatews.NewCallBack(func(msg *gatews.UpdateMsg) {
// // 	var trade gatews.SpotTradeMsg
// // 	if err := json.Unmarshal(msg.Result, &trade); err != nil {
// // 		log.Printf("trade Unmarshal err:%s", err.Error())
// // 	}
// // 	// log.Println(trade)
// // 	log.Printf("TRADE %+v", trade)
// // })

// // create callback functions for receive messages
// // callOrder := gatews.NewCallBack(func(msg *gatews.UpdateMsg) {
// // 	// parse the message to struct we need
// // 	var order []gatews.SpotOrderMsg
// // 	if err := json.Unmarshal(msg.Result, &order); err != nil {
// // 		log.Printf("order Unmarshal err:%s", err.Error())
// // 	}
// // 	// log.Println(order)
// // 	log.Printf("ORDER %+v", order)
// // })

// // callOrderBookUpdate := gatews.NewCallBack(func(msg *gatews.UpdateMsg) {
// // 	var OrderBookUpdate gatews.SpotUpdateDepthMsg
// // 	if err := json.Unmarshal(msg.Result, &OrderBookUpdate); err != nil {
// // 		log.Printf("OrderBookUpdate Unmarshal err:%s", err.Error())
// // 	}
// // 	// log.Println(depth)
// // 	log.Printf("OrderBookUpdate %+v", OrderBookUpdate)
// // })

// // GateIoWSConnService.SetCallBack(gatews.ChannelSpotOrderBookUpdate, callOrderBookUpdate)

// // // second, after set callback function, subscribe to ChannelSpotOrderBook any channel you are interested into
// // if err := GateIoWSConnService.Subscribe(gatews.ChannelSpotOrderBook, []string{"BTC_USDT", "ETH_USDT", "XRP_USDT", "LTC_BTC", "XRP_BTC"}); err != nil {
// // 	log.Printf("Subscribe err:%s", err.Error())
// // 	return
// // }

// // first, we need set callback function
// // GateIoWSConnService.SetCallBack(gatews.ChannelSpotPublicTrade, callTrade)
// // GateIoWSConnService.SetCallBack(gatews.ChannelSpotOrder, callOrder)
// // GateIoWSConnService.SetCallBack(gatews.ChannelSpotOrderBook, callOrderBook)
// // GateIoWSConnService.SetCallBack(gatews.ChannelSpotOrderBookUpdate, callOrderBookUpdate)

// // // second, after set callback function, subscribe to ChannelSpotPublicTrade any channel you are interested into
// // if err := GateIoWSConnService.Subscribe(gatews.ChannelSpotPublicTrade, []string{"BTC_USDT", "ETH_USDT", "XRP_USDT", "LTC_BTC", "XRP_BTC"}); err != nil {
// // 	log.Printf("Subscribe err:%s", err.Error())
// // 	return
// // }

// // second, after set callback function, subscribe to ChannelSpotBookTicker any channel you are interested into
// // if err := GateIoWSConnService.Subscribe(gatews.ChannelSpotBookTicker, []string{"BTC_USDT", "ETH_USDT", "XRP_USDT", "LTC_BTC", "XRP_BTC"}); err != nil {
// // 	log.Printf("Subscribe err:%s", err.Error())
// // 	return
// // }

// // // second, after set callback function, subscribe to ChannelSpotOrderBook any channel you are interested into
// // if err := GateIoWSConnService.Subscribe(gatews.ChannelSpotOrderBook, []string{"BTC_USDT", "ETH_USDT", "XRP_USDT", "LTC_BTC", "XRP_BTC"}); err != nil {
// // 	log.Printf("Subscribe err:%s", err.Error())
// // 	return
// // }

// // }
