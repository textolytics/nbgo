package strategy

// Ticks
// Depths
// Orders

type strategyService struct {
	eventGateIoWsApi EventGateIoWsApi
}

func NewStrategyService(eventGateIoWsApi EventGateIoWsApi) StrategyService {

	return &strategyService{
		eventGateIoWsApi,
	}
}

func (s *strategyService) GetTick(instrumentID string) (Tick []string, err error) {

	return s.eventGateIoWsApi.GetTicks(Instrument1, Instrument2, Instrument3)
}

func (s *strategyService) GetDepth(instrumentID string) (Depth []string, err error) {

	return s.eventGateIoWsApi.GetDepths(Instrument1, Instrument2, Instrument3)
}

func (s *strategyService) SendOrder(instrumentID string) (err error) {

	return s.eventGateIoWsApi.SendOrders(Instrument1, Instrument2, Instrument3)
}

// import (
// 	"errors"
// 	"time"

// 	errs "github.com/pkg/errors"
// 	"github.com/teris-io/shortid"
// 	"gopkg.in/dealancer/validate.v2"
// )

// var (
// 	ErrRedirectNotFound = errors.New("Redirect Not Found")
// 	ErrRedirectInvalid  = errors.New("Redirect Invalid")
// )

// type redirectService struct {
// 	redirectRepo RedirectRepository
// }

// func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
// 	return &redirectService{
// 		redirectRepo,
// 	}
// }

// func (r *redirectService) Find(code string) (*Redirect, error) {
// 	return r.redirectRepo.Find(code)
// }

// func (r *redirectService) Store(redirect *Redirect) error {
// 	if err := validate.Validate(redirect); err != nil {
// 		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
// 	}
// 	redirect.Code = shortid.MustGenerate()
// 	redirect.CreatedAt = time.Now().UTC().Unix()
// 	return r.redirectRepo.Store(redirect)
// }
