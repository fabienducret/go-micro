package main

import (
	"log"
	"mailer-service/adapters"
	"mailer-service/server"
)

func main() {
	log.Println("Starting mail service")
	s := server.NewServer(adapters.NewMailhogRepository())

	s.Listen()
}
