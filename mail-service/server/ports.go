package server

import "mailer-service/entities"

type MailRepository interface {
	SendSMTPMessage(entities.Message) error
}
