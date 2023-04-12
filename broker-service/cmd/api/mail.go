package main

import (
	"broker/ports"
)

func SendMail(mr ports.Mailer, from, to, subject, message string) (*jsonResponse, error) {
	msg := ports.Mail{
		From:    from,
		To:      to,
		Subject: subject,
		Message: message,
	}

	err := mr.Send(msg)
	if err != nil {
		return nil, err
	}

	payload := &jsonResponse{
		Error:   false,
		Message: "Message sent to " + to,
	}

	return payload, nil
}
