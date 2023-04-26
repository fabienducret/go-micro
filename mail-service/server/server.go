package server

import (
	"fmt"
	"log"
	"mailer-service/ports"
	"net"
	"net/rpc"
)

const port = "5001"

type server struct {
	MailerRepository ports.MailRepository
}

type Payload struct {
	From    string
	To      string
	Subject string
	Message string
}

func NewServer(mr ports.MailRepository) *server {
	s := new(server)
	s.MailerRepository = mr

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

func (s *server) SendMail(payload Payload, resp *string) error {
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
