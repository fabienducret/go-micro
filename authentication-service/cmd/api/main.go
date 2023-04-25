package main

import (
	"authentication/repositories"
	"fmt"
	"log"
	"net"
	"net/rpc"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const port = "5001"

type Config struct {
}

func main() {
	log.Println("Starting authentication service")

	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres")
	}

	server := new(Server)
	server.UserRepository = repositories.NewPostgresRepository(conn)
	server.LoggerRepository = repositories.NewLoggerRepository()

	start(server)
}

func start(server *Server) {
	err := rpc.Register(server)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting RPC server on port ", port)
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		rpcConn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}
