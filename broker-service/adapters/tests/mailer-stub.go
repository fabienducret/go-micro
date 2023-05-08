package tests

import (
	"broker/ports"
	"fmt"
)

type mailerStub struct{}

func NewMailerStub() *mailerStub {
	return &mailerStub{}
}

func (r *mailerStub) Send(mail ports.Mail) (string, error) {
	return fmt.Sprintf("Message sent to %s", mail.To), nil
}
