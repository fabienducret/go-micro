package adapters

import (
	"authentication/ports"
	"net/rpc"
)

const loggerServiceAddress = "logger-service:5001"

type payload struct {
	Name string
	Data string
}

type logger struct{}

func NewLogger() *logger {
	return &logger{}
}

func (r *logger) Log(toLog ports.Log) error {
	client, err := rpc.Dial("tcp", loggerServiceAddress)
	if err != nil {
		return err
	}

	err = client.Call("Server.LogInfo", payload(toLog), nil)
	if err != nil {
		return err
	}

	return nil
}
