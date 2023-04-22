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

func (a authenticateServiceRepository) AuthenticateWith(credentials ports.Credentials) (string, error) {
	client, err := rpc.Dial("tcp", authenticateServiceAddress)
	if err != nil {
		return "", err
	}

	var replyFromCall string
	err = client.Call("Server.Authenticate", authPayload(credentials), &replyFromCall)
	if err != nil {
		return "", err
	}

	return replyFromCall, nil
}
