package tests

import (
	"broker/entities"
	"errors"
	"fmt"
)

type MailerStub struct {
	WithError bool
}

func (m MailerStub) Send(mail entities.Mail) (string, error) {
	if m.WithError {
		return "", errors.New("error in mailer stub")
	}

	return fmt.Sprintf("Message sent to %s", mail.To), nil
}
