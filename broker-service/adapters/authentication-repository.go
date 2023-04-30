package adapters

import (
	"broker/ports"
	"net/rpc"
)

type authenticationRepository struct {
	addr string
}

func NewAuthenticationRepository(addr string) *authenticationRepository {
	return &authenticationRepository{
		addr: addr,
	}
}

func (a authenticationRepository) AuthenticateWith(credentials ports.Credentials) (*ports.Identity, error) {
	client, err := rpc.Dial("tcp", a.addr)
	if err != nil {
		return nil, err
	}

	var identity *ports.Identity
	err = client.Call("Server.Authenticate", credentials, &identity)
	if err != nil {
		return nil, err
	}

	return identity, nil
}
