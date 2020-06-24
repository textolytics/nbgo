package dw
//EventSerializer to 
type EventSerializer interface {
	Decode(event []byte) (*Event, error)
	Encode(event *Event) ([]byte, error)
}
