package operator

type Event struct {
	Type      string
	Body      string
	Timestamp int64
}

// type Event struct {
// 	Type      string `json:"code" bson:"code" msgpack:"code"`
// 	Body      string `json:"url" bson:"url" msgpack:"url" validate:"empty=false & format=url`
// 	Timestamp int64  `json:"created_at" bson:"created_at" msgpack:"created_at"`
// }
