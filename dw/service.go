package dw
// EventService to store all events
type EventService interface {
	Find(code string) (*Event, error)
	Store(event *Event) error
}
