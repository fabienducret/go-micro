package main

import (
	"broker/ports"
)

func SendMail(mr ports.Mailer, payload ports.MailPayload) (*jsonResponse, error) {
	mail := ports.Mail{
		From:    payload.From,
		To:      payload.To,
		Subject: payload.Subject,
		Message: payload.Message,
	}

	err := mr.Send(mail)
	if err != nil {
		return nil, err
	}

	return mailSentPayload(mail.To), nil
}

func mailSentPayload(to string) *jsonResponse {
	return &jsonResponse{
		Error:   false,
		Message: "Message sent to " + to,
	}
}
