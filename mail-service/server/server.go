package server

import (
	"fmt"
	"log"
	"mailer-service/ports"
	"net"
	"net/rpc"
)

const port = "5001"

type Server struct {
	MailerRepository ports.MailRepository
}

type Payload struct {
	From    string
	To      string
	Subject string
	Message string
}

func NewServer(mr ports.MailRepository) *Server {
	s := new(Server)
	s.MailerRepository = mr

	return s
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

func (s *Server) SendMail(payload Payload, resp *string) error {
	msg := ports.Message{
		From:    payload.From,
		To:      payload.To,
		Subject: payload.Subject,
		Data:    payload.Message,
	}

	err := s.MailerRepository.SendSMTPMessage(msg)
	if err != nil {
		return err
	}

	*resp = "Message sent to " + payload.To

	return nil
}
