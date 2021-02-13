package oanda

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

//Client; Instruments,Stream>> Schema ;  Test
//Test Account, Instrument JSON Schema
var accountID string

func TestGetPricing(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	accountID = os.Getenv("OANDA_ACCOUNT_ID")
	got := accountID
	if got != "101-004-3748257-001" {
		t.Errorf("accountID = %s; want 101-004-3748257-001 got: ", got)
	} else {
		t.Logf("accountID = %s; want 101-004-3748257-001 got: ", got)
	}
}

func BenchmarkGetCredentials(b *testing.B) {
	for i := 0; i < b.N; i++ {
		accountID = os.Getenv("OANDA_ACCOUNT_ID")
	}
}

// func TestOandaInstruments(oanda goanda.OandaConnection, accountID string) (oandaInstrumentsList goanda.AccountInstruments) {
// 	oandaInstrumentsList = oanda.GetAccountInstruments(accountID)
// 	spew.Dump("%+v\n", oandaInstrumentsList.Instruments)

// 	fmt.Println("-----------------------------------------------------")

// 	fmt.Println(oandaInstrumentsList.Instruments)
// 	return oandaInstrumentsList
// }
