package main

import (
	"broker/adapters"
	"broker/config"
	"broker/handlers"
	"broker/server"
)

func main() {
	c := config.Get()

	auth := adapters.NewAuthentication(c.AuthenticationServiceAddress, c.AuthenticationServiceMethod)
	logger := adapters.NewLogger(c.LoggerServiceAddress, c.LoggerServiceMethod)
	mailer := adapters.NewMailer(c.MailerServiceAddress, c.MailerServiceMethod)

	s := server.New(server.Handlers{
		Broker:  handlers.Broker,
		Message: handlers.MessageFactory(auth, logger, mailer),
	})

	if err := s.Run(c.Port); err != nil {
		panic(err)
	}
}
