package server

import (
	"broker/adapters"
	"broker/config"
	"broker/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func routes(c config.Config) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/", handlers.Broker)

	mux.Post("/message", message(c))

	return mux
}

func message(c config.Config) func(w http.ResponseWriter, r *http.Request) {
	auth := adapters.NewAuthentication(c.AuthenticationServiceAddress, c.AuthenticationServiceMethod)
	logger := adapters.NewLogger(c.LoggerServiceAddress, c.LoggerServiceMethod)
	mailer := adapters.NewMailer(c.MailerServiceAddress, c.MailerServiceMethod)

	return handlers.MessageFactory(auth, logger, mailer)
}
