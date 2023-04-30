package main

import (
	"log"
	"log-service/adapters"
	"log-service/db"
	"log-service/server"
	"os"
)

func main() {
	log.Println("Starting logger service")

	client, err := db.Connect(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Panic(err)
	}
	defer db.Disconnect(client)

	s := server.NewServer(adapters.NewMongoRepository(client))

	s.Listen()
}
