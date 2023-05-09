package tests

import (
	"broker/ports"
	"fmt"
)

type MailerStub struct {
	Error error
}

func (r MailerStub) Send(mail ports.Mail) (string, error) {
	return fmt.Sprintf("Message sent to %s", mail.To), r.Error
}
