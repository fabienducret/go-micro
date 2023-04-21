package repositories

import (
	"authentication/ports"
	"net/rpc"
)

type payload struct {
	Name string
	Data string
}

type loggerRepository struct{}

func NewLoggerRepository() *loggerRepository {
	return &loggerRepository{}
}

func (r *loggerRepository) Log(toLog ports.Log) error {
	client, err := rpc.Dial("tcp", "logger-service:5001")
	if err != nil {
		return err
	}

	err = client.Call("RPCServer.LogInfo", payload(toLog), nil)
	if err != nil {
		return err
	}

	return nil
}
