package adapters

import (
	"broker/ports"
	"net/rpc"
)

type logger struct {
	addr string
}

func NewLogger(addr string) *logger {
	return &logger{
		addr: addr,
	}
}

func (l *logger) Log(toLog ports.Log) (string, error) {
	client, err := rpc.Dial("tcp", l.addr)
	if err != nil {
		return "", err
	}

	var reply string
	err = client.Call("Server.LogInfo", toLog, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}
