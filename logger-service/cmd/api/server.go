package main

import (
	"log"
	"log-service/data"
	"time"
)

type Server struct {
	models data.Models
}

type Payload struct {
	Name string
	Data string
}

func (r *Server) LogInfo(payload Payload, resp *string) error {
	err := r.models.LogEntry.Insert(data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Log handled for:" + payload.Name

	return nil
}
