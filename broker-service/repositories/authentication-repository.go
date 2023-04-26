package repositories

import (
	"broker/ports"
	"net/rpc"
)

const authenticationAddress = "authentication-service:5001"

type authenticationRepository struct{}

func NewAuthenticationRepository() *authenticationRepository {
	return &authenticationRepository{}
}

func (a authenticationRepository) AuthenticateWith(credentials ports.Credentials) (*ports.Identity, error) {
	client, err := rpc.Dial("tcp", authenticationAddress)
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
