package operator

type EventSerializer interface {
	Decode(input []byte) (*Event, error)
	Encode(input *Event) ([]byte, error)
}
