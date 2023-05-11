package server

import "authentication/entities"

type Logger interface {
	Log(entities.Log) error
}

type UserRepository interface {
	GetByEmail(email string) (*entities.User, error)
	PasswordMatches(u entities.User, plainText string) (bool, error)
}
