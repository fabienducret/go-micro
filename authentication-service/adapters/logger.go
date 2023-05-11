package adapters

import (
	"authentication/entities"
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

func (r *logger) Log(toLog entities.Log) error {
	client, err := rpc.Dial("tcp", r.addr)
	if err != nil {
		return err
	}

	err = client.Call("Server.LogInfo", toLog, nil)
	if err != nil {
		return err
	}

	return nil
}
