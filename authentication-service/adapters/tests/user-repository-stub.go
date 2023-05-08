package tests

import (
	"authentication/ports"
)

type userRepositoryStub struct {
}

func NewUserRepositoryStub() *userRepositoryStub {
	return &userRepositoryStub{}
}

func (r *userRepositoryStub) GetByEmail(email string) (*ports.User, error) {
	user := ports.User{
		Email:     "test@gmail.com",
		FirstName: "Homer",
		LastName:  "Simpson",
		Password:  "password",
	}

	return &user, nil
}

func (p *userRepositoryStub) PasswordMatches(u ports.User, plainText string) (bool, error) {
	if u.Password != plainText {
		return false, nil
	}

	return true, nil
}
