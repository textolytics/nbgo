package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awoldes/goanda"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func getCredentials() (key string, accountID string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key = os.Getenv("OANDA_API_KEY")
	accountID = os.Getenv("OANDA_ACCOUNT_ID")
	return key, accountID
}

func getClient(key string, accountID string) (oanda *goanda.OandaConnection) {
	oanda = goanda.NewConnection(accountID, key, false)
	return oanda
}

func getOandaInstrumentsDetails(oanda goanda.OandaConnection, accountID string) (oandaInstrumentsDetails goanda.AccountInstruments) {
	oandaInstrumentsDetails = oanda.GetAccountInstruments(accountID)
	spew.Dump("%+v\n", oandaInstrumentsDetails)
	return oandaInstrumentsDetails
}

func getOandaInstrumentsList(oandaInstrumentsDetails goanda.AccountInstruments) (oandaInstrumentsList []string) {
	for _, v := range oandaInstrumentsDetails.Instruments {
		oandaInstrumentsList = append(oandaInstrumentsList, v.Name)
		fmt.Println(v.Name)
	}
	return oandaInstrumentsList
}

//Client; Instruments,Stream>> Schema ;  Test
func getPricing(oandaInstrumentsList []string, oanda *goanda.OandaConnection) (orderResponse goanda.Pricings) {
	orderResponse = oanda.GetPricingForInstruments(oandaInstrumentsList)
	return orderResponse
}

func getStreaming(oandaInstrumentsList []string, oanda *goanda.OandaConnection) (orderResponse goanda.Pricings) {
	orderResponse = oanda.GetPricingForInstruments(oandaInstrumentsList)
	return orderResponse
}

func main() {
	token, account := getCredentials()
	oandaClient := getClient(token, account)
	oandaInstrumentsDetails := getOandaInstrumentsDetails(*oandaClient, account)
	oandaInstrumentsList := getOandaInstrumentsList(oandaInstrumentsDetails)
	orderResponse := getPricing(oandaInstrumentsList, oandaClient)
	spew.Dump("%+v\n", orderResponse)
}
