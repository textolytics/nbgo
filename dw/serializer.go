package dw
//EventSerializer to 
type EventSerializer interface {
	Decode(input []byte) (*Event, error)
	Encode(input *Event) ([]byte, error)
}
