package config

import "os"

type Config struct {
	Port     string
	MongoUrl string
}

func Get() Config {
	return Config{
		Port:     os.Getenv("PORT"),
		MongoUrl: os.Getenv("MONGO_URL"),
	}
}
