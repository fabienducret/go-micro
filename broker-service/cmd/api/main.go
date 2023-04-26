package main

import (
	"broker/ports"
	"broker/repositories"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type App struct {
	AuthenticationRepository ports.Authentication
	LoggerRepository         ports.Logger
	MailerRepository         ports.Mailer
}

func main() {
	app := App{
		AuthenticationRepository: repositories.NewAuthenticationRepository(),
		LoggerRepository:         repositories.NewLoggerRepository(),
		MailerRepository:         repositories.NewMailerRepository(),
	}

	log.Printf("Starting broker service on port %s\n", webPort)

	start(app)
}

func start(app App) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
