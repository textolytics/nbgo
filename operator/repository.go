package operator

// DwRepository - DataWarehousing
type DwRepository interface {
	Find(data string) (*Event, error)
	Store(data *Event) error
}
