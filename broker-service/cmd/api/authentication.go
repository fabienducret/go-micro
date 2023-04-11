package main

import (
	"broker/ports"
	"errors"
)

func Authenticate(asr ports.AuthenticationService, email string, password string) (*jsonResponse, error) {
	creds := ports.Credentials{
		Email:    email,
		Password: password,
	}

	response, err := asr.AuthenticateWith(creds)
	if err != nil {
		return nil, err
	}

	if response.Error {
		return nil, errors.New("error in authenticate response")
	}

	return authenticatedPayload(response.Data), nil
}

func authenticatedPayload(data any) *jsonResponse {
	return &jsonResponse{
		Error:   false,
		Message: "Authenticated !",
		Data:    data,
	}
}
