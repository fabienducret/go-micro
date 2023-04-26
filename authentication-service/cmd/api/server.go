package main

import (
	"authentication/ports"
	"errors"
	"fmt"
)

type Server struct {
	UserRepository   ports.UserRepository
	LoggerRepository ports.Logger
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
	err = s.LoggerRepository.Log(toLog)
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
