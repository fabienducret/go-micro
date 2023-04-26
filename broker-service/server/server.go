package server

import (
	"broker/ports"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type server struct {
	AuthenticationRepository ports.Authentication
	LoggerRepository         ports.Logger
	MailerRepository         ports.Mailer
}

func NewServer(ar ports.Authentication, lr ports.Logger, mr ports.Mailer) *server {
	return &server{
		AuthenticationRepository: ar,
		LoggerRepository:         lr,
		MailerRepository:         mr,
	}
}

func (s *server) Start() {
	log.Printf("Starting broker service on port %s\n", webPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: s.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
