package main

import (
	"broker/ports"
	"errors"
	"net/http"
)

func (app *App) Broker(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	})
}

func (app *App) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload ports.RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.handleAuthenticate(w, requestPayload)
	case "log":
		app.handleLog(w, requestPayload)
	case "mail":
		app.handleMail(w, requestPayload)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *App) handleAuthenticate(w http.ResponseWriter, requestPayload ports.RequestPayload) {
	asr := app.Container.AuthenticationServiceRepository
	reply, err := asr.AuthenticateWith(ports.Credentials(requestPayload.Auth))
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: "Authenticated !",
		Data:    reply,
	})
}

func (app *App) handleLog(w http.ResponseWriter, requestPayload ports.RequestPayload) {
	lr := app.Container.LoggerRepository
	reply, err := lr.Log(ports.Log(requestPayload.Log))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: reply,
	})
}

func (app *App) handleMail(w http.ResponseWriter, requestPayload ports.RequestPayload) {
	mr := app.Container.MailerRepository
	reply, err := mr.Send(ports.Mail(requestPayload.Mail))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: reply,
	})
}
