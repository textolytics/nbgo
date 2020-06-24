package dw

// EventRepository - DataWarehousing
type EventRepository interface {
	Find(data string) (*Event, error)
	Store(data *Event) error
}
