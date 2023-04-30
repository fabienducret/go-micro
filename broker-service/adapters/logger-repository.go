package adapters

import (
	"broker/ports"
	"net/rpc"
)

type loggerRepository struct {
	addr string
}

func NewLoggerRepository(addr string) *loggerRepository {
	return &loggerRepository{
		addr: addr,
	}
}

func (l *loggerRepository) Log(toLog ports.Log) (string, error) {
	client, err := rpc.Dial("tcp", l.addr)
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
