package repositories

import (
	"broker/ports"
	"net/rpc"
)

type loggerRepository struct {
}

type RPCPayload struct {
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

	payload := RPCPayload{
		Name: toLog.Name,
		Data: toLog.Data,
	}

	var result string
	err = client.Call("RPCServer.LogInfo", payload, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}
