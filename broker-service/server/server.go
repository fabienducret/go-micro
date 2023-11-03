package server

import (
	"fmt"
	"log"
	"net/http"
)

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

func (s *server) Run(port string) {
	log.Printf("Starting broker service on port %s\n", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.Routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
