package main

import (
	"log"
	"log-service/ports"
)

type Server struct {
	LogRepository ports.LogRepository
}

type Payload struct {
	Name string
	Data string
}

func (r *Server) LogInfo(payload Payload, resp *string) error {
	err := r.LogRepository.Insert(ports.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Log handled for:" + payload.Name

	return nil
}
