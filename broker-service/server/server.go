package server

import (
	"broker/config"
	"fmt"
	"log"
	"net/http"
)

func RunWith(c config.Config) error {
	port := c.Port
	log.Printf("Starting broker service on port %s\n", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: initHandlersWith(c),
	}

	return server.ListenAndServe()
}
