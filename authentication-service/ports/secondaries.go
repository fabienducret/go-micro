package ports

type Log struct {
	Name string
	Data string
}

type User struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

type Logger interface {
	Log(Log) error
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	PasswordMatches(u User, plainText string) (bool, error)
}
