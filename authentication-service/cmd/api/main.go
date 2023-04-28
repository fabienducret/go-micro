package main

import (
	"authentication/adapters"
	"authentication/db"
	"authentication/server"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	log.Println("Starting authentication service")

	conn := db.Connect()
	if conn == nil {
		log.Panic("Can't connect to Postgres")
	}

	s := server.NewServer(
		adapters.NewPostgresRepository(conn),
		adapters.NewLoggerRepository(),
	)

	s.Listen()
}
