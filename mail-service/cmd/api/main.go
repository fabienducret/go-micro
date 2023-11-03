package main

import (
	"log"
	"mailer-service/adapters"
	"mailer-service/config"
	"mailer-service/server"
)

func main() {
	log.Println("Starting mail service")

	c := config.Get()
	s := server.New(adapters.NewMailhogRepository(c))

	s.Listen()
}
