package main

import (
	"log"
	"log-service/adapters"
	"log-service/config"
	"log-service/db"
	"log-service/server"
)

func main() {
	log.Println("Starting logger service")
	c := config.Get()

	client, err := db.Connect(c.MongoUrl)
	if err != nil {
		log.Panic(err)
	}
	defer db.Disconnect(client)

	s := server.NewServer(adapters.NewMongoRepository(client))

	s.Listen(c.Port)
}
