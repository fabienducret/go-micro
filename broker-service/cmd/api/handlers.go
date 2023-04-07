package main

import (
	"broker/ports"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmissionFactory(asr ports.AuthenticationService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestPayload RequestPayload

		err := app.readJSON(w, r, &requestPayload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		switch requestPayload.Action {
		case "auth":
			payload, err := app.Authenticate(asr, requestPayload.Auth.Email, requestPayload.Auth.Password)
			if err != nil {
				app.errorJSON(w, err, http.StatusUnauthorized)
				return
			}

			app.writeJSON(w, http.StatusAccepted, payload)
		default:
			app.errorJSON(w, errors.New("unknown action"))
		}
	}
}
