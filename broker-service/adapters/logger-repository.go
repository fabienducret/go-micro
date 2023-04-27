package adapters

import (
	"broker/ports"
	"net/rpc"
)

const loggerServiceAddress = "logger-service:5001"

type loggerRepository struct{}

func NewLoggerRepository() *loggerRepository {
	return &loggerRepository{}
}

func (l *loggerRepository) Log(toLog ports.Log) (string, error) {
	client, err := rpc.Dial("tcp", loggerServiceAddress)
	if err != nil {
		return "", err
	}

	var replyFromCall string
	err = client.Call("Server.LogInfo", toLog, &replyFromCall)
	if err != nil {
		return "", err
	}

	return replyFromCall, nil
}
