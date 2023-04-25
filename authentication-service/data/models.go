package data

import (
	"authentication/ports"
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const dbTimeout = time.Second * 3

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"-"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostgresRepository struct {
	Conn *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: db,
	}
}

// GetByEmail returns one user by email
func (r *PostgresRepository) GetByEmail(email string) (*ports.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where email = $1`

	var fromDB User
	row := r.Conn.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&fromDB.ID,
		&fromDB.Email,
		&fromDB.FirstName,
		&fromDB.LastName,
		&fromDB.Password,
		&fromDB.Active,
		&fromDB.CreatedAt,
		&fromDB.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	user := ports.User{
		Email:     fromDB.Email,
		FirstName: fromDB.FirstName,
		LastName:  fromDB.LastName,
		Password:  fromDB.Password,
	}

	return &user, nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (p *PostgresRepository) PasswordMatches(u ports.User, plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
