package server

import (
	"broker/ports"
	"errors"
	"log"
	"net/http"
)

func (s *server) Broker(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, Payload{
		Error:   false,
		Message: "Hit the broker",
	})
}

func (s *server) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var request ports.RequestPayload

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

func (s *server) handleAuthenticate(w http.ResponseWriter, payload ports.AuthPayload) {
	reply, err := s.authentication.AuthenticateWith(ports.Credentials(payload))
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

func (s *server) handleLog(w http.ResponseWriter, payload ports.LogPayload) {
	reply, err := s.logger.Log(ports.Log(payload))
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

func (s *server) handleMail(w http.ResponseWriter, payload ports.MailPayload) {
	reply, err := s.mailer.Send(ports.Mail(payload))
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
