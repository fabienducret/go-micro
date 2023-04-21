package main

import (
	"broker/ports"
)

func Authenticate(asr ports.AuthenticationService, payload ports.AuthPayload) (*jsonResponse, error) {
	creds := ports.Credentials(payload)

	response, err := asr.AuthenticateWith(creds)
	if err != nil {
		return nil, err
	}

	return authenticatedPayload(response), nil
}

func authenticatedPayload(response string) *jsonResponse {
	return &jsonResponse{
		Error:   false,
		Message: "Authenticated !",
		Data:    response,
	}
}
