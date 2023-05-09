package server

import (
	"broker/ports"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type server struct {
	authentication ports.Authentication
	logger         ports.Logger
	mailer         ports.Mailer
}

func NewServer(
	auth ports.Authentication,
	logger ports.Logger,
	mailer ports.Mailer,
) *server {
	return &server{
		authentication: auth,
		logger:         logger,
		mailer:         mailer,
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
