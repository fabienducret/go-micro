package tests

import (
	"authentication/entities"
	"errors"
)

type LoggerStub struct {
	WithError bool
}

func (l LoggerStub) Log(toLog entities.Log) error {
	if l.WithError {
		return errors.New("error in logger stub")
	}

	return nil
}
