package main

import (
	"log"
	"os"

	"github.com/awoldes/goanda"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

//Client ; Instruments,Stream>> Schema ;  Test
func getPricing() (oanda *goanda.OandaConnection, accountID string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := os.Getenv("OANDA_API_KEY")
	accountID = os.Getenv("OANDA_ACCOUNT_ID")
	oanda = goanda.NewConnection(accountID, key, false)
	oandaInstruments(*oanda, accountID)

	instruments := []string{"AUD_USD", "EUR_NZD"}
	orderResponse := oanda.GetPricingForInstruments(instruments)
	spew.Dump("%+v\n", orderResponse)
	// fmt.Println(orderResponse)
	return oanda, accountID
}

func oandaInstruments(oanda goanda.OandaConnection, accountID string) (oandaInstrumentsList goanda.AccountInstruments) {

	oandaInstrumentsList = oanda.GetAccountInstruments(accountID)
	spew.Dump("%+v\n", oandaInstrumentsList.Instruments)
	// fmt.Println((oanda.GetAccountInstruments).Instruments())

	return oandaInstrumentsList
}

func main() {
	getPricing()

}
