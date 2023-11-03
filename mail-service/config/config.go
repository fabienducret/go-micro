package config

import "os"

type Config struct {
	Port           string
	MailPort       string
	MailDomain     string
	MailHost       string
	MailUsername   string
	MailPassword   string
	MailEncryption string
	FromName       string
	FromAddress    string
}

func Get() Config {
	return Config{
		Port:           os.Getenv("PORT"),
		MailPort:       os.Getenv("MAIL_PORT"),
		MailDomain:     os.Getenv("MAIL_DOMAIN"),
		MailHost:       os.Getenv("MAIL_HOST"),
		MailUsername:   os.Getenv("MAIL_USERNAME"),
		MailPassword:   os.Getenv("MAIL_PASSWORD"),
		MailEncryption: os.Getenv("MAIL_ENCRYPTION"),
		FromName:       os.Getenv("FROM_NAME"),
		FromAddress:    os.Getenv("FROM_ADDRESS"),
	}
}
