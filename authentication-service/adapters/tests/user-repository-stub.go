package tests

import (
	"authentication/ports"
)

type UserRepositoryStub struct{}

func (r UserRepositoryStub) GetByEmail(email string) (*ports.User, error) {
	user := ports.User{
		Email:     "test@gmail.com",
		FirstName: "Homer",
		LastName:  "Simpson",
		Password:  "password",
	}

	return &user, nil
}

func (p UserRepositoryStub) PasswordMatches(u ports.User, plainText string) (bool, error) {
	if u.Password != plainText {
		return false, nil
	}

	return true, nil
}
