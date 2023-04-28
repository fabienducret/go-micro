package main

import (
	"log"
	"log-service/adapters"
	"log-service/data"
	"log-service/server"
)

func main() {
	log.Println("Starting logger service")

	client, err := data.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}
	defer data.DisconnectClient(client)

	s := server.NewServer(adapters.NewMongoRepository(client))

	s.Listen()
}
