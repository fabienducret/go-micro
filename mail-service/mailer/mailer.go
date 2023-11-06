package mailer

import "mailer-service/entities"

type MailRepository interface {
	SendSMTPMessage(entities.Message) error
}

type Mailer struct {
	MailerRepository MailRepository
}

type Payload struct {
	From    string
	To      string
	Subject string
	Message string
}

func New(mr MailRepository) *Mailer {
	m := new(Mailer)
	m.MailerRepository = mr

	return m
}

func (s *Mailer) SendMail(payload Payload, resp *string) error {
	msg := entities.Message{
		From:    payload.From,
		To:      payload.To,
		Subject: payload.Subject,
		Data:    payload.Message,
	}

	err := s.MailerRepository.SendSMTPMessage(msg)
	if err != nil {
		return err
	}

	*resp = "Message sent to " + payload.To

	return nil
}
