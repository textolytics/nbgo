package main

import (
	"net/url"
	"strings"

	"github.com/gookit/config"
)

var GATEIO_API_KEY, _ = config.String("exchanges.gateio_api_key")
var GATEIO_API_SECRET, _ = config.String("exchanges.gateio_api_secret")

type RunConfig struct {
	ApiKey     string
	ApiSecret  string
	BaseUrl    string
	UseTestNet bool
}

func NewRunConfig(apiKey string, apiSecret string, hostUsed *string) (*RunConfig, error) {
	config := &RunConfig{
		ApiKey:     GATEIO_API_KEY,
		ApiSecret:  GATEIO_API_SECRET,
		UseTestNet: false,
		BaseUrl:    *hostUsed,
	}
	if hostUsed == nil || *hostUsed == "" {
		config.BaseUrl = "https://api.gateio.ws/api/v4"
	}
	if !strings.HasPrefix(config.BaseUrl, "http") {
		config.BaseUrl = "https://" + config.BaseUrl
	}
	if !strings.HasSuffix(config.BaseUrl, "/api/v4") {
		config.BaseUrl += "/api/v4"
	}
	parsedUrl, err := url.Parse(config.BaseUrl)
	if err != nil {
		return nil, err
	}
	if parsedUrl.Host == "fx-api-testnet.gateio.ws" {
		config.UseTestNet = true
	}
	return config, nil
}
