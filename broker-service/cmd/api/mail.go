package main

import (
	"broker/ports"
)

func SendMail(mr ports.Mailer, payload ports.MailPayload) (*jsonResponse, error) {
	mail := ports.Mail(payload)

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
