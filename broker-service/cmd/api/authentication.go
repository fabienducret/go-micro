package main

import (
	"broker/ports"
	"encoding/json"
)

func (app *Config) Authenticate(asr ports.AuthenticationService, email string, password string) (jsonResponse, error) {
	var payload jsonResponse

	body, err := asr.AuthenticateWith(email, password)
	if err != nil {
		return payload, err
	}
	defer body.Close()

	var jsonFromService jsonResponse
	err = json.NewDecoder(body).Decode(&jsonFromService)
	if err != nil {
		return payload, err
	}

	if jsonFromService.Error {
		return payload, err
	}

	payload = authenticatedPayload(jsonFromService)

	return payload, nil
}

func authenticatedPayload(jsonFromService jsonResponse) jsonResponse {
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated !"
	payload.Data = jsonFromService.Data

	return payload
}
