package main

import (
	"log"
	"log-service/adapters"
	"log-service/config"
	"log-service/db"
	"log-service/listener"
	"log-service/logger"
)

func main() {
	log.Println("Starting logger service")
	c := config.Get()

	client, err := db.Connect(c.MongoUrl)
	if err != nil {
		log.Panic(err)
	}
	defer db.Disconnect(client)

	logger := logger.New(adapters.NewMongoRepository(client))
	l := listener.New(logger)

	if err := l.Listen(c.Port); err != nil {
		panic(err)
	}
}
