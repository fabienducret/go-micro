package repositories

import (
	"context"
	"log"
	"log-service/ports"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type MongoRepository struct {
	Client *mongo.Client
}

func NewMongoRepository(db *mongo.Client) *MongoRepository {
	return &MongoRepository{
		Client: db,
	}
}

func (r *MongoRepository) Insert(entry ports.LogEntry) error {
	collection := r.Client.Database("logs").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting into logs: ", err)
		return err
	}

	return nil
}
