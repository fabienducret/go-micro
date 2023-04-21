package main

import (
	"log"
	"log-service/data"
	"time"
)

type RPCServer struct {
	models data.Models
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	err := r.models.LogEntry.Insert(data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Processed payload via RPC:" + payload.Name

	return nil
}
