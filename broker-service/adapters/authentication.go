package adapters

import (
	"broker/entities"
	"net/rpc"
)

type authentication struct {
	addr   string
	method string
}

func NewAuthentication(addr, method string) *authentication {
	return &authentication{
		addr,
		method,
	}
}

func (a authentication) AuthenticateWith(credentials entities.Credentials) (*entities.Identity, error) {
	client, err := rpc.Dial("tcp", a.addr)
	if err != nil {
		return nil, err
	}

	var identity *entities.Identity
	err = client.Call(a.method, credentials, &identity)
	if err != nil {
		return nil, err
	}

	return identity, nil
}
