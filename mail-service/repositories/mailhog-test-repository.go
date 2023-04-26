package repositories

import (
	"mailer-service/ports"
)

type mailhogTestRepository struct {
}

func NewMailhogTestRepository() *mailhogTestRepository {
	return &mailhogTestRepository{}
}

func (r *mailhogTestRepository) SendSMTPMessage(msg ports.Message) error {
	return nil
}
