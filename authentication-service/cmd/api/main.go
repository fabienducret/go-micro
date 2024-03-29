package main

import (
	"authentication/adapters"
	"authentication/authentication"
	"authentication/config"
	"authentication/db"
	"authentication/listener"
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
		panic("Can't connect to Postgres")
	}

	a := authentication.New(
		adapters.NewPostgresRepository(conn),
		adapters.NewLogger(c.LoggerServiceAddress, c.LoggerServiceMethod),
	)
	l := listener.New(a)

	if err := l.Listen(c.Port); err != nil {
		panic(err)
	}
}
