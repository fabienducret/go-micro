package main

import (
	"broker/ports"
	"errors"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload ports.RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		asr := app.Container.AuthenticationServiceRepository
		payload, err := Authenticate(asr, requestPayload.Auth.Email, requestPayload.Auth.Password)
		if err != nil {
			app.errorJSON(w, err, http.StatusUnauthorized)
			return
		}

		app.writeJSON(w, http.StatusAccepted, payload)
	case "log":
		lr := app.Container.LoggerRepository
		payload, err := Log(lr, requestPayload.Log.Name, requestPayload.Log.Data)
		if err != nil {
			app.errorJSON(w, err)
		}

		app.writeJSON(w, http.StatusAccepted, payload)
	case "mail":
		mr := app.Container.MailerRepository
		payload, err := SendMail(
			mr,
			requestPayload.Mail.From,
			requestPayload.Mail.To,
			requestPayload.Mail.Subject,
			requestPayload.Mail.Message,
		)
		if err != nil {
			app.errorJSON(w, err)
		}

		app.writeJSON(w, http.StatusAccepted, payload)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}
