package dw

import (
	"errors"
	errs "github.com/pkg/errors"
	validate "gopkg.in/go-playground/validator.v9"
)

var (
	// ErrEventNotFound to log not found
	ErrEventNotFound = errors.New("Event Not Found")
	// ErrEventInvalid to log invalid
	ErrEventInvalid = errors.New("Event Invalid")
)

type eventService struct {
	eventRepo EventRepository
}

//NewEventService to process event
func NewEventService(eventRepo EventRepository) EventService {
	return &eventService{
		eventRepo,
	}
}

func (r *eventService) Find(code string) (*Event, error) {
	return r.eventRepo.Find(code)
}

func (r *eventService) Store(event *Event) error {
	if err := validate.New().Struct(event); err != nil {
		return errs.Wrap(ErrEventInvalid, "service.Event.Store")
	}
	errs.Wrap(ErrEventInvalid, "service.Event.Store")
	// event.Body = shortid.MustGenerate()
	// event.Timestamp = time.Now().UTC().Unix()

	return r.eventRepo.Store(event)
}

// func (r *redirectService) Store(redirect *Redirect) error {
// 	if err := validate.Validate(redirect); err != nil {
// 		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
// 	}
// 	redirect.Code = shortid.MustGenerate()
// 	redirect.CreatedAt = time.Now().UTC().Unix()
// 	return r.redirectRepo.Store(redirect)
// }
