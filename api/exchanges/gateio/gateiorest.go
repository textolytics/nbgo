package gateio

// import (
// 	"context"

// 	"github.com/antihax/optional"
// 	"github.com/gateio/gateapi-go/v6"
// 	"github.com/shopspring/decimal"
// )

// func SpotDemo(config *RunConfig) {
// 	client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// 	// Setting host is optional. It defaults to https://api.gateio.ws/api/v4
// 	client.ChangeBasePath(config.BaseUrl)
// 	ctx := context.WithValue(context.Background(), gateapi.ContextGateAPIV4, gateapi.GateAPIV4{
// 		Key:    config.ApiKey,
// 		Secret: config.ApiSecret,
// 	})

// 	currencyPair := "GT_USDT"
// 	currency := "USDT"
// 	cp, _, err := client.SpotApi.GetCurrencyPair(ctx, currencyPair)
// 	if err != nil {
// 		panicGateError(err)
// 	}
// 	logger.Printf("testing against currency pair: %s\n", cp.Id)
// 	minAmount := cp.MinBaseAmount

// 	tickers, _, err := client.SpotApi.ListTickers(ctx, &gateapi.ListTickersOpts{CurrencyPair: optional.NewString(cp.Id)})
// 	if err != nil {
// 		panicGateError(err)
// 	}
// 	lastPrice := tickers[0].Last

// 	// better avoid using float, take the following decimal library for example
// 	// `go get github.com/shopspring/decimal`
// 	orderAmount := decimal.RequireFromString(minAmount).Mul(decimal.NewFromInt32(2))

// 	balance, _, err := client.SpotApi.ListSpotAccounts(ctx, &gateapi.ListSpotAccountsOpts{Currency: optional.NewString(currency)})
// 	if err != nil {
// 		panicGateError(err)
// 	}
// 	if decimal.RequireFromString(balance[0].Available).Cmp(orderAmount) < 0 {
// 		logger.Fatal("balance not enough")
// 	}

// 	newOrder := gateapi.Order{
// 		Text:         "t-my-custom-id", // optional custom order ID
// 		CurrencyPair: cp.Id,
// 		Type:         "limit",
// 		Account:      "spot", // create spot order. set to "margin" if creating margin orders
// 		Side:         "buy",
// 		Amount:       orderAmount.String(),
// 		Price:        lastPrice, // use last price
// 		TimeInForce:  "gtc",
// 		AutoBorrow:   false,
// 	}
// 	logger.Printf("place a spot %s order in %s with amount %s and price %s\n", newOrder.Side, newOrder.CurrencyPair, newOrder.Amount, newOrder.Price)
// 	createdOrder, _, err := client.SpotApi.CreateOrder(ctx, newOrder)
// 	if err != nil {
// 		panicGateError(err)
// 	}
// 	logger.Printf("order created with ID: %s, status: %s\n", createdOrder.Id, createdOrder.Status)
// 	if createdOrder.Status == "open" {
// 		order, _, err := client.SpotApi.GetOrder(ctx, createdOrder.Id, createdOrder.CurrencyPair)
// 		if err != nil {
// 			panicGateError(err)
// 		}
// 		logger.Printf("order %s filled: %s, left: %s\n", order.Id, order.FilledTotal, order.Left)
// 		result, _, err := client.SpotApi.CancelOrder(ctx, createdOrder.Id, createdOrder.CurrencyPair)
// 		if err != nil {
// 			panicGateError(err)
// 		}
// 		if result.Status == "cancelled" {
// 			logger.Printf("order %s cancelled\n", createdOrder.Id)
// 		}
// 	} else {
// 		// order finished
// 		trades, _, err := client.SpotApi.ListMyTrades(ctx, createdOrder.CurrencyPair,
// 			&gateapi.ListMyTradesOpts{OrderId: optional.NewString(createdOrder.Id)})
// 		if err != nil {
// 			panicGateError(err)
// 		}
// 		for _, t := range trades {
// 			logger.Printf("order %s filled %s with price: %s\n", t.OrderId, t.Amount, t.Price)
// 		}
// 	}
// }

// // package gateio

// // import (
// // 	"net/url"
// // 	"strings"
// // 	"github.com/gateio/gateapi-go/v5"
// // 	"github.com/gookit/config"
// // )

// // var GATEIOREST_API_KEY, _ = config.String("exchanges.gateio.gateio_api_key")
// // var GATEIOREST_API_SECRET, _ = config.String("exchanges.gateio.gateio_api_secret")

// // // var GATEIORESTCONFIGFUNCRETURN = gateiorest.NewConfiguration()

// // // type gateiorestConfig gateiorest.Configuration

// // // type gateiorestClient gateiorest.APIClient

// // // var gateiorestConfig = &gateiorest.Configuration{
// // // 	Key:    GATEIOREST_API_KEY,
// // // 	Secret: GATEIOREST_API_SECRET,
// // // }

// // // var gateiorestClient = &gateiorest.APIClient{}

// // type RunConfig struct {
// // 	ApiKey     string
// // 	ApiSecret  string
// // 	BaseUrl    string
// // 	UseTestNet bool
// // }

// // func NewRunConfig(apiKey string, apiSecret string, hostUsed *string) (*RunConfig, error) {
// // 	config := &RunConfig{
// // 		ApiKey:     GATEIOREST_API_KEY,
// // 		ApiSecret:  GATEIOREST_API_SECRET,
// // 		UseTestNet: false,
// // 		BaseUrl:    *hostUsed,
// // 	}
// // 	if hostUsed == nil || *hostUsed == "" {
// // 		config.BaseUrl = "https://api.gateio.ws/api/v4"
// // 	}
// // 	if !strings.HasPrefix(config.BaseUrl, "http") {
// // 		config.BaseUrl = "https://" + config.BaseUrl
// // 	}
// // 	if !strings.HasSuffix(config.BaseUrl, "/api/v4") {
// // 		config.BaseUrl += "/api/v4"
// // 	}
// // 	parsedUrl, err := url.Parse(config.BaseUrl)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	if parsedUrl.Host == "fx-api-testnet.gateio.ws" {
// // 		config.UseTestNet = true
// // 	}
// // 	return config, nil
// // }
// // var client = gateapi.APIClient

// // func GateIoConnect(config *RunConfig)
// // 	client = gateapi.NewAPIClient(gateapi.NewConfiguration())
// // 	// Setting host is optional. It defaults to https://api.gateio.ws/api/v4
// // 	client.ChangeBasePath(config.BaseUrl)
// // 	ctx := context.WithValue(context.Background(), gateapi.ContextGateAPIV4, gateapi.GateAPIV4{
// // 		Key:    config.ApiKey,
// // 		Secret: config.ApiSecret,
// // })

// // func main (){

// // 	GateIoConnect()

// // }

// // // var gateiorestClient = gateiorest.APIClient{
// // // 	cfg: gateiorestConfig,
// // // }

// // // // // var GATEIORESTCONFIGTYPE = gateiorest.NewConfiguration().Key

// // // var NewGateioRestConfig = gateiorest.NewConfiguration()

// // // func makeGateiorestConfig() *gateiorestConfig {

// // // 	gateiorestConfig := &gateiorestConfig{}
// // // 	return gateiorestConfig
// // // }
// // // var gateiorest.SetKeySecret =

// // // func GateIoRestClient(GATEIOREST_API_KEY string, GATEIOREST_API_SECRET string) gateiorestClient {

// // // 	client := gateiorest.NewAPIClient(cfg)

// // // 	// gateiorest.SetKeySecret(GATEIOREST_API_KEY, GATEIOREST_API_SECRET)
// // // 	return client
// // // }

// // // var GATEIORESTCONFIG = gateiorest.NewConfiguration()
// // // {
// // // 	Key:    GATEIO_API_KEY,
// // // 	Secret: GATEIO_API_SECRET,
// // // 	// MaxRetryConn:  10, // default value is math.MaxInt64, set it when needs
// // // 	// SkipTlsVerify: false,
// // // }

// // // {
// // // 	Key:    GATEIOREST_API_KEY,
// // // 	Secret: GATEIOREST_API_SECRET,
// // // }

// // // type GateIoRestClient interface {
// // // 	Connect(gateiorest.APIClient) gateiorest.APIClient
// // // 	Disconnect()
// // // 	Sttaus()
// // // }

// // // type GateIoRestApiService gateiorest.SpotApiService

// // // type GateIoRestEvent interface {
// // // 	Get(gateiorest.SpotApiService) gateiorest.APIResponse
// // // 	Post(gateiorest.SpotApiService) gateiorest.APIResponse
// // // }
