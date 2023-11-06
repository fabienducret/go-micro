package main

import (
	"log"
	"mailer-service/adapters"
	"mailer-service/config"
	"mailer-service/listener"
	"mailer-service/mailer"
)

func main() {
	log.Println("Starting mail service")
	c := config.Get()

	mailer := mailer.New(adapters.NewMailhogRepository(c))
	l := listener.New(mailer)

	err := l.Listen(c.Port)
	if err != nil {
		panic(err)
	}
}
