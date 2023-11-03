package server

import (
	"broker/config"
	"fmt"
	"log"
	"net/http"
)

func RunWith(c config.Config) {
	port := c.Port
	log.Printf("Starting broker service on port %s\n", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes(c),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
