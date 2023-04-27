package main

import (
	"context"
	"log"
	"log-service/adapters"
	"log-service/data"
	"log-service/server"
	"time"
)

func main() {
	log.Println("Starting logger service")

	client, err := data.ConnectToMongo()
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

	s := server.NewServer(adapters.NewMongoRepository(client))

	s.Listen()
}
