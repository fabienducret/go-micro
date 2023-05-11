package tests

import (
	"errors"
	"mailer-service/entities"
)

type MailRepositoryStub struct {
	WithError bool
}

func (m MailRepositoryStub) SendSMTPMessage(msg entities.Message) error {
	if m.WithError {
		return errors.New("error in mail repository stub")
	}

	return nil
}
