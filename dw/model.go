package main

// //Event to work with events
// type Event struct {
// }

//Route to handle venue events
type Route interface {
}

//Event to work with events
type Event struct {
	Type      string `json:"code" bson:"code" msgpack:"code"`
	Body      string `json:"url" bson:"url" msgpack:"url" validate:"empty=false & format=url`
	Timestamp int64  `json:"created_at" bson:"created_at" msgpack:"created_at"`
}

// type EventRepository struct {
// }
