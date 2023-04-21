package ports

type Credentials struct {
	Email    string
	Password string
}

type Log struct {
	Name string
	Data string
}

type Mail struct {
	From    string
	To      string
	Subject string
	Message string
}

type AuthenticationService interface {
	AuthenticateWith(Credentials) (string, error)
}

type Logger interface {
	Log(Log) (string, error)
}

type Mailer interface {
	Send(Mail) (string, error)
}
