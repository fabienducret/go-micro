package adapters

import (
	"broker/ports"
	"fmt"
)

type testMailer struct{}

func NewTestMailer() *testMailer {
	return &testMailer{}
}

func (r *testMailer) Send(mail ports.Mail) (string, error) {
	return fmt.Sprintf("Message sent to %s", mail.To), nil
}
