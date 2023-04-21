package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net"
	"net/rpc"
	"time"
)

const port = "5001"

type Config struct {
	Models data.Models
}

func main() {
	log.Println("Starting logger service")

	client, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	app.listen()
}

func (app *Config) listen() {
	server := new(Server)
	server.models = app.Models

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
