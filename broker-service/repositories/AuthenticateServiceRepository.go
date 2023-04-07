package repositories

import (
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

type authenticateServiceRepository struct{}

func NewAuthenticateServiceRepository() *authenticateServiceRepository {
	return &authenticateServiceRepository{}
}

func (a authenticateServiceRepository) AuthenticateWith(email string, password string) (io.ReadCloser, error) {
	var authPayload authPayload
	authPayload.Email = email
	authPayload.Password = password

	jsonData, _ := json.MarshalIndent(authPayload, "", "\t")

	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("invalid credentials")
	} else if response.StatusCode != http.StatusAccepted {
		return nil, errors.New("error calling auth service")
	}

	return response.Body, nil
}
