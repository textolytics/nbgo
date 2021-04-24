package nboanda

import (
	"fmt"
	"log"
	"os"

	"github.com/awoldes/goanda"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func getCredentials() (key string, accountID string) {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/textolytics/nbgo/api/nboanda/.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	keyName := ("OANDA_API_KEY")
	key = os.Getenv("OANDA_API_KEY")

	accountIDName := ("OANDA_ACCOUNT_ID")
	accountID = os.Getenv("OANDA_ACCOUNT_ID")
	os.Setenv(keyName, key)
	os.Setenv(accountIDName, accountID)

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

//GetPricing; Instruments,Stream>> Schema ;  Test
func getPricing(oandaInstrumentsList []string, oanda *goanda.OandaConnection) (orderResponse goanda.Pricings) {
	orderResponse = oanda.GetPricingForInstruments(oandaInstrumentsList)
	return orderResponse
}

//GetOandaPricing Instruments,Stream>> Schema ;  Test
func GetOandaPricing() (orderResponse goanda.Pricings) {
	token, account := getCredentials()
	oandaClient := getClient(token, account)
	// getStreamingClient := getStreamingClient(token, account)
	oandaInstrumentsDetails := getOandaInstrumentsDetails(*oandaClient, account)
	oandaInstrumentsList := getOandaInstrumentsList(oandaInstrumentsDetails)
	orderResponse = getPricing(oandaInstrumentsList, oandaClient)
	spew.Dump("%+v\n", orderResponse)
	// streamingResponse = getStreaming(oandaInstrumentsList, getStreamingClient)
	// spew.Dump("%+v\n", getStreamingClient)
	return orderResponse
}

// func main() {

// }
