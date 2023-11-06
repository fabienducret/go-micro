package authentication

import (
	"authentication/entities"
	"errors"
	"fmt"
)

type Logger interface {
	Log(entities.Log) error
}

type UserRepository interface {
	GetByEmail(email string) (*entities.User, error)
	PasswordMatches(u entities.User, plainText string) (bool, error)
}

type Authentication struct {
	UserRepository UserRepository
	Logger         Logger
}

type Payload struct {
	Email    string
	Password string
}

type Identity struct {
	Email     string
	FirstName string
	LastName  string
}

func New(ur UserRepository, l Logger) *Authentication {
	s := new(Authentication)
	s.UserRepository = ur
	s.Logger = l

	return s
}

func (s *Authentication) Authenticate(payload Payload, reply *Identity) error {
	user, err := s.UserRepository.GetByEmail(payload.Email)
	if err != nil {
		return errors.New("invalid credentials")
	}

	valid, err := s.UserRepository.PasswordMatches(*user, payload.Password)
	if err != nil || !valid {
		return errors.New("invalid credentials")
	}

	toLog := entities.Log{
		Name: "authentication",
		Data: fmt.Sprintf("%s logged in", user.Email),
	}
	err = s.Logger.Log(toLog)
	if err != nil {
		return err
	}

	*reply = Identity{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return nil
}
