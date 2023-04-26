package repositories

import (
	"authentication/ports"
)

type PostgresTestRepository struct {
}

func NewPostgresTestRepository() *PostgresTestRepository {
	return &PostgresTestRepository{}
}

func (r *PostgresTestRepository) GetByEmail(email string) (*ports.User, error) {
	user := ports.User{
		Email:     "test@gmail.com",
		FirstName: "Homer",
		LastName:  "Simpson",
		Password:  "password",
	}

	return &user, nil
}

func (p *PostgresTestRepository) PasswordMatches(u ports.User, plainText string) (bool, error) {
	if u.Password != plainText {
		return false, nil
	}

	return true, nil
}
