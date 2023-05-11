package server

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

func (s *server) Broker(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, Payload{
		Error:   false,
		Message: "Hit the broker",
	})
}

func (s *server) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var request RequestPayload

	err := readJSON(w, r, &request)
	if err != nil {
		errorJSON(w, err)
		return
	}

	switch request.Action {
	case "auth":
		s.handleAuthenticate(w, request.Auth)
	case "log":
		s.handleLog(w, request.Log)
	case "mail":
		s.handleMail(w, request.Mail)
	default:
		errorJSON(w, errors.New("unknown action"))
	}
}

func (s *server) handleAuthenticate(w http.ResponseWriter, payload AuthPayload) {
	reply, err := s.authentication.AuthenticateWith(entities.Credentials(payload))
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

func (s *server) handleLog(w http.ResponseWriter, payload LogPayload) {
	reply, err := s.logger.Log(entities.Log(payload))
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

func (s *server) handleMail(w http.ResponseWriter, payload MailPayload) {
	reply, err := s.mailer.Send(entities.Mail(payload))
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
