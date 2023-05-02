package adapters

import (
	"broker/ports"
	"net/rpc"
)

type mailer struct {
	addr string
}

func NewMailer(addr string) *mailer {
	return &mailer{
		addr: addr,
	}
}

func (r *mailer) Send(mail ports.Mail) (string, error) {
	client, err := rpc.Dial("tcp", r.addr)
	if err != nil {
		return "", err
	}

	var reply string
	err = client.Call("Server.SendMail", mail, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}
