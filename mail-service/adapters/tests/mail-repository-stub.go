package tests

import (
	"errors"
	"mailer-service/ports"
)

type MailRepositoryStub struct {
	WithError bool
}

func (m MailRepositoryStub) SendSMTPMessage(msg ports.Message) error {
	if m.WithError {
		return errors.New("error in mail repository stub")
	}

	return nil
}
