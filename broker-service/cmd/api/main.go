package main

import (
	"broker/repositories"
	"broker/server"
)

func main() {
	s := server.NewServer(
		repositories.NewAuthenticationRepository(),
		repositories.NewLoggerRepository(),
		repositories.NewMailerRepository(),
	)

	s.Start()
}
