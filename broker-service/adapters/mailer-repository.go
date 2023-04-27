package adapters

import (
	"broker/ports"
	"net/rpc"
)

const mailServiceAddress = "mail-service:5001"

type mailerRepository struct{}

func NewMailerRepository() *mailerRepository {
	return &mailerRepository{}
}

func (r *mailerRepository) Send(mail ports.Mail) (string, error) {
	client, err := rpc.Dial("tcp", mailServiceAddress)
	if err != nil {
		return "", err
	}

	var replyFromCall string
	err = client.Call("Server.SendMail", mail, &replyFromCall)
	if err != nil {
		return "", err
	}

	return replyFromCall, nil
}
