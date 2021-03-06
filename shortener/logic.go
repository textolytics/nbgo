package shortener

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	validate "gopkg.in/go-playground/validator.v9"
)

var (
	//ErrRedirectNotFound sdsdsd
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	//ErrRedirectInvalid sdsds
	ErrRedirectInvalid = errors.New("Redirect Invalid")
)

type redirectService struct {
	redirectRepo RedirectRepository
}

//NewRedirectService (redirectRepo RedirectRepository) RedirectService
func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepo,
	}
}

func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	if err := validate.New().Struct(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}

// func (r *redirectService) Store(redirect *Redirect) error {
// 	if err := validate.Validate(redirect); err != nil {
// 		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
// 	}
// 	redirect.Code = shortid.MustGenerate()
// 	redirect.CreatedAt = time.Now().UTC().Unix()
// 	return r.redirectRepo.Store(redirect)
// }
