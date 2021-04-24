package dw

// DataWarehouseDataFrame
type DataWarehouseDataFrame struct {
	Destination string `json:"destination" bson:"destination" msgpack:"destination"`
	Data        string `json:"data" bson:"data" msgpack:"data"`
}

// // Redirect struct {
// type Redirect struct {
// 	Code      string `json:"code" bson:"code" msgpack:"code"`
// 	URL       string `json:"url" bson:"url" msgpack:"url" validate:"empty=false & format=url`
// 	CreatedAt int64  `json:"created_at" bson:"created_at" msgpack:"created_at"`
// }
