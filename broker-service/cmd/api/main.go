package main

import (
	"broker/repositories"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	app := Config{}
	asr := repositories.NewAuthenticateServiceRepository()

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(asr),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
