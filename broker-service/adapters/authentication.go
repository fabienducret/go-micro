package adapters

import (
	"broker/entities"
	"net/rpc"
)

type authentication struct {
	addr string
}

func NewAuthentication(addr string) *authentication {
	return &authentication{
		addr: addr,
	}
}

func (a authentication) AuthenticateWith(credentials entities.Credentials) (*entities.Identity, error) {
	client, err := rpc.Dial("tcp", a.addr)
	if err != nil {
		return nil, err
	}

	var identity *entities.Identity
	err = client.Call("Server.Authenticate", credentials, &identity)
	if err != nil {
		return nil, err
	}

	return identity, nil
}
