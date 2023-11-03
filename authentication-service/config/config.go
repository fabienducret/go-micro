package config

import "os"

type Config struct {
	Port                 string
	LoggerServiceAddress string
	DatabaseDsn          string
}

func Get() Config {
	return Config{
		Port:                 os.Getenv("PORT"),
		LoggerServiceAddress: os.Getenv("LOGGER_SERVICE_ADDRESS"),
		DatabaseDsn:          os.Getenv("DSN"),
	}
}
