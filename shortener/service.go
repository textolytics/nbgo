package shortener

//RedirectService dwindows 10 built dools
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
