package adapters

import (
	"authentication/entities"
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

func (r *logger) Log(toLog entities.Log) error {
	client, err := rpc.Dial("tcp", r.addr)
	if err != nil {
		return err
	}

	err = client.Call(r.method, toLog, nil)
	if err != nil {
		return err
	}

	return nil
}
