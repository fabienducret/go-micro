package ports

type Credentials struct {
	Email    string
	Password string
}

type Log struct {
	Name string
	Data string
}

type AuthenticateResponse struct {
	Error bool
	Data  any
}

type AuthenticationService interface {
	AuthenticateWith(Credentials) (*AuthenticateResponse, error)
}

type Logger interface {
	Log(Log) error
}
