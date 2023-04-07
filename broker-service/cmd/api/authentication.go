package main

import (
	"broker/ports"
	"errors"
)

func (app *Config) Authenticate(asr ports.AuthenticationService, email string, password string) (*jsonResponse, error) {
	response, err := asr.AuthenticateWith(email, password)
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
