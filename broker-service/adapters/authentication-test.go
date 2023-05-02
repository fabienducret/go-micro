package adapters

import (
	"broker/ports"
	"errors"
)

type testAuthentication struct{}

func NewTestAuthentication() *testAuthentication {
	return &testAuthentication{}
}

func (a testAuthentication) AuthenticateWith(credentials ports.Credentials) (*ports.Identity, error) {
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
