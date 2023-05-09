package tests

import (
	"authentication/ports"
	"errors"
)

type LoggerStub struct {
	WithError bool
}

func (l LoggerStub) Log(toLog ports.Log) error {
	if l.WithError {
		return errors.New("error in logger stub")
	}

	return nil
}
