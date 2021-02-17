package dw

// DataframeRepository interface {
type DataframeRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}

// RedirectRepository interface {
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
