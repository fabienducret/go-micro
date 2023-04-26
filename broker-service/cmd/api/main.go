package main

import (
	"broker/repositories"
	"broker/server"
)

func main() {
	server.NewServer(
		repositories.NewAuthenticationRepository(),
		repositories.NewLoggerRepository(),
		repositories.NewMailerRepository(),
	).Start()
}
