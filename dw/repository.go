package main

// EventRepository - DataWarehousing
type EventRepository interface {
	Find(event string) (*Event, error)
	Store(event *Event) error
}
