package main

import (
	"authentication/data"
	"authentication/repositories"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/rpc"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const rpcPort = "5001"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	// Connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	app.startServer()
}

func (app *Config) startServer() {
	rpcServer := new(RPCServer)
	rpcServer.Models = app.Models
	rpcServer.LoggerRepository = repositories.NewLoggerRepository()

	_ = rpc.Register(rpcServer)

	log.Println("Starting RPC server on port ", rpcPort)
	listen, _ := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}
