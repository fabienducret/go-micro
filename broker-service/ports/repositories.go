package ports

type AuthenticateResponse struct {
	Error bool
	Data  any
}

type AuthenticationService interface {
	AuthenticateWith(email string, password string) (*AuthenticateResponse, error)
}
