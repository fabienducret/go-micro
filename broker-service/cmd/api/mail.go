package main

import (
	"broker/ports"
)

func SendMail(mr ports.Mailer, payload ports.MailPayload) (*jsonResponse, error) {
	mail := ports.Mail(payload)

	reply, err := mr.Send(mail)
	if err != nil {
		return nil, err
	}

	return mailSentPayload(reply), nil
}

func mailSentPayload(reply string) *jsonResponse {
	return &jsonResponse{
		Error:   false,
		Message: reply,
	}
}
