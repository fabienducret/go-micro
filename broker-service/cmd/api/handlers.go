package main

import (
	"broker/ports"
	"errors"
	"net/http"
)

func (app *App) Broker(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, Payload{
		Error:   false,
		Message: "Hit the broker",
	})
}

func (app *App) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var request ports.RequestPayload

	err := readJSON(w, r, &request)
	if err != nil {
		errorJSON(w, err)
		return
	}

	switch request.Action {
	case "auth":
		app.handleAuthenticate(w, request.Auth)
	case "log":
		app.handleLog(w, request.Log)
	case "mail":
		app.handleMail(w, request.Mail)
	default:
		errorJSON(w, errors.New("unknown action"))
	}
}

func (app *App) handleAuthenticate(w http.ResponseWriter, payload ports.AuthPayload) {
	reply, err := app.AuthenticationRepository.AuthenticateWith(ports.Credentials(payload))
	if err != nil {
		errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	writeJSON(w, http.StatusAccepted, Payload{
		Error:   false,
		Message: "Authenticated !",
		Data:    reply,
	})
}

func (app *App) handleLog(w http.ResponseWriter, payload ports.LogPayload) {
	reply, err := app.LoggerRepository.Log(ports.Log(payload))
	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, http.StatusAccepted, Payload{
		Error:   false,
		Message: reply,
	})
}

func (app *App) handleMail(w http.ResponseWriter, payload ports.MailPayload) {
	reply, err := app.MailerRepository.Send(ports.Mail(payload))
	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, http.StatusAccepted, Payload{
		Error:   false,
		Message: reply,
	})
}
