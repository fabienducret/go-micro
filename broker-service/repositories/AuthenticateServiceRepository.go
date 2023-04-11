package repositories

import (
	"broker/ports"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type authPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const authenticateUrl = "http://authentication-service/authenticate"

type authenticateServiceRepository struct{}

func NewAuthenticateServiceRepository() *authenticateServiceRepository {
	return &authenticateServiceRepository{}
}

func (a authenticateServiceRepository) AuthenticateWith(credentials ports.Credentials) (*ports.AuthenticateResponse, error) {
	toSend := formatCredentialsRequest(credentials)

	request, err := http.NewRequest("POST", authenticateUrl, toSend)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("invalid credentials")
	} else if response.StatusCode != http.StatusAccepted {
		return nil, errors.New("error calling auth service")
	}

	return parseResponse(response.Body)
}

func formatCredentialsRequest(credentials ports.Credentials) *bytes.Buffer {
	authPayload := authPayload{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	jsonData, _ := json.MarshalIndent(authPayload, "", "\t")

	return bytes.NewBuffer(jsonData)
}

func parseResponse(body io.ReadCloser) (*ports.AuthenticateResponse, error) {
	authenticateResponse := &ports.AuthenticateResponse{}
	err := json.NewDecoder(body).Decode(&authenticateResponse)
	if err != nil {
		return nil, err
	}

	return authenticateResponse, nil
}
