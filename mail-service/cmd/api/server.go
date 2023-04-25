package main

import "mailer-service/ports"

type Server struct {
	MailerRepository ports.MailRepository
}

type Payload struct {
	From    string
	To      string
	Subject string
	Message string
}

func (r *Server) SendMail(payload Payload, resp *string) error {
	msg := ports.Message{
		From:    payload.From,
		To:      payload.To,
		Subject: payload.Subject,
		Data:    payload.Message,
	}

	err := r.MailerRepository.SendSMTPMessage(msg)
	if err != nil {
		return err
	}

	*resp = "Message sent to " + payload.To

	return nil
}
