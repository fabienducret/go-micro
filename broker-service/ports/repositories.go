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

type Identity struct {
	Email     string
	FirstName string
	LastName  string
}

type Authentication interface {
	AuthenticateWith(Credentials) (*Identity, error)
}

type Logger interface {
	Log(Log) (string, error)
}

type Mailer interface {
	Send(Mail) (string, error)
}
