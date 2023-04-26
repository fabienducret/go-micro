package main

import (
	"context"
	"log"
	"log-service/repositories"
	"log-service/server"
	"time"
)

func main() {
	log.Println("Starting logger service")

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

	s := server.NewServer(repositories.NewMongoRepository(client))

	s.Listen()
}
