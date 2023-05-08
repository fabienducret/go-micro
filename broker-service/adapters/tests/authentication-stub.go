package tests

import (
	"broker/ports"
	"errors"
)

type authenticationStub struct{}

func NewAuthenticationStub() *authenticationStub {
	return &authenticationStub{}
}

func (a authenticationStub) AuthenticateWith(credentials ports.Credentials) (*ports.Identity, error) {
	if credentials.Password != "verysecret" {
		return nil, errors.New("invalid password")
	}

	identity := ports.Identity{
		Email:     "homer@simpson.com",
		FirstName: "Homer",
		LastName:  "Simpson",
	}

	return &identity, nil
}
