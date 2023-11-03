package adapters

import (
	"broker/entities"
	"net/rpc"
)

type mailer struct {
	addr   string
	method string
}

func NewMailer(addr, method string) *mailer {
	return &mailer{
		addr,
		method,
	}
}

func (r *mailer) Send(mail entities.Mail) (string, error) {
	client, err := rpc.Dial("tcp", r.addr)
	if err != nil {
		return "", err
	}

	var reply string
	err = client.Call(r.method, mail, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}
