package dw
// EventService to store all events
type EventService interface {
	Find(event string) (*Event, error)
	Store(event *Event) error
}
