package main

import (
	"authentication/adapters"
	"authentication/config"
	"authentication/db"
	"authentication/server"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	log.Println("Starting authentication service")
	c := config.Get()

	conn := db.Connect(c.DatabaseDsn)
	if conn == nil {
		log.Panic("Can't connect to Postgres")
	}

	s := server.New(
		adapters.NewPostgresRepository(conn),
		adapters.NewLogger(c.LoggerServiceAddress),
	)

	s.Listen(c.Port)
}
