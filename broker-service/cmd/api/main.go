package main

import (
	"broker/adapters"
	"broker/server"
	"os"
)

func main() {
	s := server.NewServer(
		adapters.NewAuthenticationRepository(os.Getenv("AUTHENTICATION_SERVICE_ADDRESS")),
		adapters.NewLoggerRepository(os.Getenv("LOGGER_SERVICE_ADDRESS")),
		adapters.NewMailerRepository(os.Getenv("MAIL_SERVICE_ADDRESS")),
	)

	s.Start()
}
