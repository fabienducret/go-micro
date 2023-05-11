package tests

import (
	"broker/entities"
	"errors"
)

type AuthenticationStub struct{}

func (a AuthenticationStub) AuthenticateWith(credentials entities.Credentials) (*entities.Identity, error) {
	if credentials.Password != "verysecret" {
		return nil, errors.New("invalid password")
	}

	identity := entities.Identity{
		Email:     "homer@simpson.com",
		FirstName: "Homer",
		LastName:  "Simpson",
	}

	return &identity, nil
}
