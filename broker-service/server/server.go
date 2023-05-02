package server

import (
	"broker/ports"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type server struct {
	Authentication ports.Authentication
	Logger         ports.Logger
	Mailer         ports.Mailer
}

func NewServer(
	auth ports.Authentication,
	logger ports.Logger,
	mailer ports.Mailer,
) *server {
	return &server{
		Authentication: auth,
		Logger:         logger,
		Mailer:         mailer,
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
