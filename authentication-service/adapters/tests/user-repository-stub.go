package tests

import (
	"authentication/entities"
)

type UserRepositoryStub struct{}

func (r UserRepositoryStub) GetByEmail(email string) (*entities.User, error) {
	user := entities.User{
		Email:     "test@gmail.com",
		FirstName: "Homer",
		LastName:  "Simpson",
		Password:  "password",
	}

	return &user, nil
}

func (p UserRepositoryStub) PasswordMatches(u entities.User, plainText string) (bool, error) {
	if u.Password != plainText {
		return false, nil
	}

	return true, nil
}
