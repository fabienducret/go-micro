package repositories

import (
	"authentication/ports"
)

type userTestRepository struct {
}

func NewUserTestRepository() *userTestRepository {
	return &userTestRepository{}
}

func (r *userTestRepository) GetByEmail(email string) (*ports.User, error) {
	user := ports.User{
		Email:     "test@gmail.com",
		FirstName: "Homer",
		LastName:  "Simpson",
		Password:  "password",
	}

	return &user, nil
}

func (p *userTestRepository) PasswordMatches(u ports.User, plainText string) (bool, error) {
	if u.Password != plainText {
		return false, nil
	}

	return true, nil
}
