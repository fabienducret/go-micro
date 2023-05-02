package server

import (
	"authentication/ports"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const port = "5001"

type Server struct {
	UserRepository ports.UserRepository
	Logger         ports.Logger
}

type Payload struct {
	Email    string
	Password string
}

type Identity struct {
	Email     string
	FirstName string
	LastName  string
}

func NewServer(ur ports.UserRepository, l ports.Logger) *Server {
	s := new(Server)
	s.UserRepository = ur
	s.Logger = l

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

func (s *Server) Authenticate(payload Payload, reply *Identity) error {
	user, err := s.UserRepository.GetByEmail(payload.Email)
	if err != nil {
		return errors.New("invalid credentials")
	}

	valid, err := s.UserRepository.PasswordMatches(*user, payload.Password)
	if err != nil || !valid {
		return errors.New("invalid credentials")
	}

	toLog := ports.Log{
		Name: "authentication",
		Data: fmt.Sprintf("%s logged in", user.Email),
	}
	err = s.Logger.Log(toLog)
	if err != nil {
		return err
	}

	*reply = Identity{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return nil
}
