package main

import (
	"log"
	"mailer-service/repositories"
	"mailer-service/server"
)

func main() {
	log.Println("Starting mail service")
	s := server.NewServer(repositories.NewMailhogRepository())

	s.Listen()
}
