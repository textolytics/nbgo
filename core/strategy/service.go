package strategy

type StrategyService interface {
	GetTick(instrumentId string) (Tick []string, err error)
	GetDepth(instrumentId string) (Depth []string, err error)
	SendOrder(instrumentId string) error
}

// type RedirectService interface {
// 	Find(code string) (*Redirect, error)
// 	Store(redirect *Redirect) error
// }
