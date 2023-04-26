package server

import (
	"fmt"
	"log"
	"log-service/ports"
	"net"
	"net/rpc"
)

const port = "5001"

type server struct {
	LogRepository ports.LogRepository
}

type Payload struct {
	Name string
	Data string
}

func NewServer(lr ports.LogRepository) *server {
	s := new(server)
	s.LogRepository = lr

	return s
}

func (s *server) Listen() {
	err := rpc.Register(s)
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

func (s *server) LogInfo(payload Payload, resp *string) error {
	err := s.LogRepository.Insert(ports.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Log handled for:" + payload.Name

	return nil
}
