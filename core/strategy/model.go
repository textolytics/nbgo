package strategy

// type Redirect struct {
// 	Code      string `json:"code" bson:"code" msgpack:"code"`
// 	URL       string `json:"url" bson:"url" msgpack:"url" validate:"empty=false & format=url`
// 	CreatedAt int64  `json:"created_at" bson:"created_at" msgpack:"created_at"`
// }

type Tick string

type Depth string

type Order string

var Instrument1 = "BTC_USDT"
var Instrument2 = "ETH_USDT"
var Instrument3 = "ETH_BTC"
