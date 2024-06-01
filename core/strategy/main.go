package strategy

import (
	"flag"
	"fmt"
	"log"

	"github.com/gateio/gateapi-go/v6"
)

var logger = log.New(flag.CommandLine.Output(), "", log.LstdFlags)

func panicGateError(err error) {
	if e, ok := err.(gateapi.GateAPIError); ok {
		log.Fatal(fmt.Sprintf("Gate API error, label: %s, message: %s", e.Label, e.Message))
	}
	log.Fatal(err)
}

func main() {

	runConfig, err := NewRunConfig(GATEIO_API_KEY, GATEIO_API_SECRET, nil)
	if err != nil {
		logger.Fatal(err)
	}
	// RunConfig
	GateioWsApiClient(runConfig)

	// var key, secret, baseUrl string
	// flag.StringVar(&key, "k", "", "Gate APIv4 key")
	// flag.StringVar(&secret, "s", "", "Gate APIv4 secret")
	// flag.StringVar(&baseUrl, "u", "", "API based URL used")
	// flag.Parse()

	// usage := fmt.Sprintf("Usage: %s -k <api-key> -s <api-secret> <spot|margin|futures>", os.Args[0])

	// if key == "" || secret == "" {
	// 	logger.Println(usage)
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }
	// if flag.NArg() < 1 {
	// 	logger.Println(usage)
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	// rand.Seed(time.Now().Unix())

	// for _, demo := range flag.Args() {
	// 	switch demo {
	// 	case "spot":
	// 		SpotDemo(runConfig)
	// 	case "margin":
	// 		MarginDemo(runConfig)
	// 	case "futures":
	// 		FuturesDemo(runConfig)
	// 	default:
	// 		logger.Fatal("Invalid demo provided. Available: spot, margin or futures")
	// 	}
	// }
}
