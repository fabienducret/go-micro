package handlers

import (
	"broker/entities"
	"errors"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
	Mail   MailPayload `json:"mail,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func HandleFactory(
	auth Authentication,
	logger Logger,
	mailer Mailer,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request RequestPayload

		err := readJSON(w, r, &request)
		if err != nil {
			errorJSON(w, err)
			return
		}

		switch request.Action {
		case "auth":
			handleAuthenticate(w, auth, request.Auth)
		case "log":
			handleLog(w, logger, request.Log)
		case "mail":
			handleMail(w, mailer, request.Mail)
		default:
			errorJSON(w, errors.New("unknown action"))
		}
	}
}

func handleAuthenticate(w http.ResponseWriter, auth Authentication, payload AuthPayload) {
	reply, err := auth.AuthenticateWith(entities.Credentials(payload))
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

func handleLog(w http.ResponseWriter, logger Logger, payload LogPayload) {
	reply, err := logger.Log(entities.Log(payload))
	if err != nil {
		log.Println(err)
		errorJSON(w, errors.New("server error on logger"), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusAccepted, Payload{
		Error:   false,
		Message: reply,
	})
}

func handleMail(w http.ResponseWriter, mailer Mailer, payload MailPayload) {
	reply, err := mailer.Send(entities.Mail(payload))
	if err != nil {
		log.Println(err)
		errorJSON(w, errors.New("server error on mail"), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusAccepted, Payload{
		Error:   false,
		Message: reply,
	})
}
