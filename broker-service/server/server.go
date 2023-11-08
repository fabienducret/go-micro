package server

import (
	"fmt"
	"log"
	"net/http"
)

type Handlers struct {
	Broker  http.HandlerFunc
	Message http.HandlerFunc
}

type server struct {
	handlers Handlers
}

func New(h Handlers) *server {
	return &server{h}
}

func (s *server) Run(port string) error {
	log.Printf("Starting broker service on port %s\n", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.initHandlers(),
	}

	return server.ListenAndServe()
}
