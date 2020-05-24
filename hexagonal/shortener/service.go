package shortener

//RedirectService ujoiuo
type RedirectService interface {
	Find(code string) (*Redirect, error) //asdasdasd
	Store(redirect *Redirect) error      // sadasdasdasd

}
