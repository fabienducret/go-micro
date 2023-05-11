package server

import (
	"fmt"
	"log"
	"log-service/entities"
	"log-service/ports"
	"net"
	"net/rpc"
)

const port = "5001"

type Server struct {
	LogRepository ports.LogRepository
}

type Payload struct {
	Name string
	Data string
}

func NewServer(lr ports.LogRepository) *Server {
	return &Server{lr}
}

func (s *Server) Listen() {
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
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

func (s *Server) LogInfo(payload Payload, resp *string) error {
	err := s.LogRepository.Insert(entities.LogEntry{
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
