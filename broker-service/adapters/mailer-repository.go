package adapters

import (
	"broker/ports"
	"net/rpc"
)

type mailerRepository struct {
	addr string
}

func NewMailerRepository(addr string) *mailerRepository {
	return &mailerRepository{
		addr: addr,
	}
}

func (r *mailerRepository) Send(mail ports.Mail) (string, error) {
	client, err := rpc.Dial("tcp", r.addr)
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
