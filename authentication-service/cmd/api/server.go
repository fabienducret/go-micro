package main

import (
	"authentication/data"
	"authentication/ports"
	"errors"
	"fmt"
)

type Server struct {
	Models           data.Models
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

func (r *Server) Authenticate(payload Payload, reply *Identity) error {
	user, err := r.Models.User.GetByEmail(payload.Email)
	if err != nil {
		return errors.New("invalid credentials")
	}

	valid, err := user.PasswordMatches(payload.Password)
	if err != nil || !valid {
		return errors.New("invalid credentials")
	}

	toLog := ports.Log{
		Name: "authentication",
		Data: fmt.Sprintf("%s logged in", user.Email),
	}
	err = r.LoggerRepository.Log(toLog)
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
