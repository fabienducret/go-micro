package ports

import "io"

type AuthenticationService interface {
	AuthenticateWith(email string, password string) (io.ReadCloser, error)
}
