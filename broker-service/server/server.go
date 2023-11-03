package server

import (
	"broker/config"
	"fmt"
	"log"
	"net/http"
)

func Run(c config.Config) {
	log.Printf("Starting broker service on port %s\n", c.Port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", c.Port),
		Handler: routes(c),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
