package adapters

import (
	"broker/ports"
	"fmt"
)

type mailerTestRepository struct{}

func NewMailerTestRepository() *mailerTestRepository {
	return &mailerTestRepository{}
}

func (r *mailerTestRepository) Send(mail ports.Mail) (string, error) {
	return fmt.Sprintf("Message sent to %s", mail.To), nil
}
