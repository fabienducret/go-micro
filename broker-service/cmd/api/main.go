package main

import (
	"broker/ports"
	"broker/repositories"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Container struct {
	AuthenticationServiceRepository ports.AuthenticationService
	LoggerRepository                ports.Logger
	MailerRepository                ports.Mailer
}

type App struct {
	Container Container
}

func main() {
	app := App{
		Container: Container{
			AuthenticationServiceRepository: repositories.NewAuthenticateServiceRepository(),
			LoggerRepository:                repositories.NewLoggerRepository(),
			MailerRepository:                repositories.NewMailerRepository(),
		},
	}

	log.Printf("Starting broker service on port %s\n", webPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	start(server)
}

func start(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
