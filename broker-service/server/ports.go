package server

import "broker/entities"

type Authentication interface {
	AuthenticateWith(entities.Credentials) (*entities.Identity, error)
}

type Logger interface {
	Log(entities.Log) (string, error)
}

type Mailer interface {
	Send(entities.Mail) (string, error)
}
