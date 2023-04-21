package repositories

import (
	"broker/ports"
	"net/rpc"
)

type mailPayload struct {
	From    string
	To      string
	Subject string
	Message string
}

type mailerRepository struct{}

func NewMailerRepository() *mailerRepository {
	return &mailerRepository{}
}

func (r *mailerRepository) Send(mail ports.Mail) (string, error) {
	client, err := rpc.Dial("tcp", "mail-service:5001")
	if err != nil {
		return "", err
	}

	var replyFromCall string
	err = client.Call("RPCServer.SendMail", mailPayload(mail), &replyFromCall)
	if err != nil {
		return "", err
	}

	return replyFromCall, nil
}
