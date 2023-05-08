package tests

import (
	"mailer-service/ports"
)

type mailRepositoryStub struct {
}

func NewMailRepositoryStub() *mailRepositoryStub {
	return &mailRepositoryStub{}
}

func (r *mailRepositoryStub) SendSMTPMessage(msg ports.Message) error {
	return nil
}
