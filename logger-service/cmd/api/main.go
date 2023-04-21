package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

const (
	rpcPort  = "5001"
	mongoUrl = "mongodb://mongo:27017"
)

type Config struct {
	Models data.Models
}

func main() {
	client, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	app.startServer()
}
