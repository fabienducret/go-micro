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

const rpcPort = "5001"

type Config struct {
	Models data.Models
}

func main() {
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

	app.startServer()
}

func (app *Config) startServer() {
	rpcServer := new(RPCServer)
	rpcServer.models = app.Models

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
