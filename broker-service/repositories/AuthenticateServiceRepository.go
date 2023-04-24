package repositories

import (
	"broker/ports"
	"net/rpc"
)

const authenticateServiceAddress = "authentication-service:5001"

type authPayload struct {
	Email    string
	Password string
}

type authenticateServiceRepository struct{}

func NewAuthenticateServiceRepository() *authenticateServiceRepository {
	return &authenticateServiceRepository{}
}

func (a authenticateServiceRepository) AuthenticateWith(credentials ports.Credentials) (*ports.Identity, error) {
	client, err := rpc.Dial("tcp", authenticateServiceAddress)
	if err != nil {
		return nil, err
	}

	var identity *ports.Identity
	err = client.Call("Server.Authenticate", authPayload(credentials), &identity)
	if err != nil {
		return nil, err
	}

	return identity, nil
}
