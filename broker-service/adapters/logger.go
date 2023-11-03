package adapters

import (
	"broker/entities"
	"net/rpc"
)

type logger struct {
	addr   string
	method string
}

func NewLogger(addr, method string) *logger {
	return &logger{
		addr,
		method,
	}
}

func (l *logger) Log(toLog entities.Log) (string, error) {
	client, err := rpc.Dial("tcp", l.addr)
	if err != nil {
		return "", err
	}

	var reply string
	err = client.Call(l.method, toLog, &reply)
	if err != nil {
		return "", err
	}

	return reply, nil
}
