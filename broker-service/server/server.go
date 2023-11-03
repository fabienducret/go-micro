package server

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type server struct {
	authentication Authentication
	logger         Logger
	mailer         Mailer
}

func NewServer(
	auth Authentication,
	logger Logger,
	mailer Mailer,
) *server {
	return &server{
		authentication: auth,
		logger:         logger,
		mailer:         mailer,
	}
}

func (s *server) Run() {
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
