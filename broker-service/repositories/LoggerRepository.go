package repositories

import (
	"broker/ports"
	"net/rpc"
)

type loggerRepository struct {
}

type payload struct {
	Name string
	Data string
}

func NewLoggerRepository() *loggerRepository {
	return &loggerRepository{}
}

func (l *loggerRepository) Log(toLog ports.Log) (string, error) {
	client, err := rpc.Dial("tcp", "logger-service:5001")
	if err != nil {
		return "", err
	}

	var replyFromCall string
	err = client.Call("RPCServer.LogInfo", payload(toLog), &replyFromCall)
	if err != nil {
		return "", err
	}

	return replyFromCall, nil
}
