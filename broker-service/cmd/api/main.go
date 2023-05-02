package main

import (
	"broker/adapters"
	"broker/server"
	"os"
)

func main() {
	s := server.NewServer(
		adapters.NewAuthentication(os.Getenv("AUTHENTICATION_SERVICE_ADDRESS")),
		adapters.NewLogger(os.Getenv("LOGGER_SERVICE_ADDRESS")),
		adapters.NewMailer(os.Getenv("MAIL_SERVICE_ADDRESS")),
	)

	s.Start()
}
