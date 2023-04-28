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
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type AuthenticationRepository interface {
	AuthenticateWith(Credentials) (*Identity, error)
}

type LoggerRepository interface {
	Log(Log) (string, error)
}

type MailerRepository interface {
	Send(Mail) (string, error)
}
