package main

import (
	"broker/adapters"
	"broker/config"
	"broker/server"
)

func main() {
	c := config.Get()

	auth := adapters.NewAuthentication(c.AuthenticationServiceAddress, c.AuthenticationServiceMethod)
	logger := adapters.NewLogger(c.LoggerServiceAddress, c.LoggerServiceMethod)
	mailer := adapters.NewMailer(c.MailerServiceAddress, c.MailerServiceMethod)

	s := server.NewServer(auth, logger, mailer)

	s.Run(c.Port)
}
