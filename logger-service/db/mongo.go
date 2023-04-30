package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(url string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(url)
	opts.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	c, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Println("Error connecting: ", err)
		return nil, err
	}

	log.Println("Starting mongodb on url", url)

	return c, nil
}

func Disconnect(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
