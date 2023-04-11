package main

import (
	"broker/ports"
	"broker/repositories"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Container struct {
	AuthenticationServiceRepository ports.AuthenticationService
	LoggerRepository                ports.Logger
}

type Config struct {
	Container Container
}

func main() {
	app := Config{
		Container{
			AuthenticationServiceRepository: repositories.NewAuthenticateServiceRepository(),
			LoggerRepository:                repositories.NewLoggerRepository(),
		},
	}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
