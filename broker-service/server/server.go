package server

import (
	"broker/ports"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type server struct {
	AuthenticationRepository ports.AuthenticationRepository
	LoggerRepository         ports.LoggerRepository
	MailerRepository         ports.MailerRepository
}

func NewServer(
	ar ports.AuthenticationRepository,
	lr ports.LoggerRepository,
	mr ports.MailerRepository,
) *server {
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
		Handler: s.Routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
