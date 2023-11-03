package config

import "os"

type Config struct {
	Port                         string
	AuthenticationServiceAddress string
	AuthenticationServiceMethod  string
	LoggerServiceAddress         string
	LoggerServiceMethod          string
	MailerServiceAddress         string
	MailerServiceMethod          string
}

func Get() Config {
	return Config{
		Port:                         os.Getenv("PORT"),
		AuthenticationServiceAddress: os.Getenv("AUTHENTICATION_SERVICE_ADDRESS"),
		AuthenticationServiceMethod:  os.Getenv("AUTHENTICATION_SERVICE_METHOD"),
		LoggerServiceAddress:         os.Getenv("LOGGER_SERVICE_ADDRESS"),
		LoggerServiceMethod:          os.Getenv("LOGGER_SERVICE_METHOD"),
		MailerServiceAddress:         os.Getenv("MAIL_SERVICE_ADDRESS"),
		MailerServiceMethod:          os.Getenv("MAIL_SERVICE_METHOD"),
	}
}
