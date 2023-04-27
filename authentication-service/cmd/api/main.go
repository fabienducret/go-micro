package main

import (
	"authentication/data"
	"authentication/repositories"
	"authentication/server"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	log.Println("Starting authentication service")

	conn := data.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres")
	}

	s := server.NewServer(
		repositories.NewPostgresRepository(conn),
		repositories.NewLoggerRepository(),
	)

	s.Listen()
}
