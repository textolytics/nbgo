package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gateio/gateapi-go/v6"
	"github.com/textolytics/nbgo/core/strategy"
)

var logger = log.New(flag.CommandLine.Output(), "", log.LstdFlags)

func panicGateError(err error) {
	if e, ok := err.(gateapi.GateAPIError); ok {
		log.Fatal(fmt.Sprintf("Gate API error, label: %s, message: %s", e.Label, e.Message))
	}
	log.Fatal(err)
}

func main() {

	// runConfig, err := NewRunConfig(GATEIO_API_KEY, GATEIO_API_SECRET, nil)
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// RunConfig
	strategy.GateioWsApiClient(&strategy.RunConfig{})
}
