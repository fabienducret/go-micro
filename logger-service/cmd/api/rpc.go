package main

import (
	"fmt"
	"log"
	"log-service/data"
	"net"
	"net/rpc"
	"time"
)

const rpcPort = "5001"

type RPCServer struct {
	models data.Models
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	err := r.models.LogEntry.Insert(data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Processed payload via RPC:" + payload.Name

	return nil
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
