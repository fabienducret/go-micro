package main

import (
	"authentication/data"
	"authentication/ports"
	"errors"
	"fmt"
)

type RPCServer struct {
	Models           data.Models
	LoggerRepository ports.Logger
}

type RPCPayload struct {
	Email    string
	Password string
}

type Response struct {
	Data any
}

func (r *RPCServer) Authenticate(payload RPCPayload, resp *string) error {
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

	*resp = fmt.Sprintf("%s logged", user.Email)

	return nil
}
