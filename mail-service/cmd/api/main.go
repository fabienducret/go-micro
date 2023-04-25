package main

import (
	"fmt"
	"log"
	"mailer-service/repositories"
	"net"
	"net/rpc"
)

const port = "5001"

func main() {
	log.Println("Starting mail service")

	server := new(Server)
	server.MailerRepository = repositories.NewMailhogRepository()

	listen(server)
}

func listen(server *Server) {
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
