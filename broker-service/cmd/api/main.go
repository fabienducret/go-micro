package main

import (
	"broker/adapters"
	"broker/server"
)

func main() {
	s := server.NewServer(
		adapters.NewAuthenticationRepository(),
		adapters.NewLoggerRepository(),
		adapters.NewMailerRepository(),
	)

	s.Start()
}
