package tests

import (
	"broker/entities"
	"errors"
	"fmt"
)

type LoggerStub struct {
	WithError bool
}

func (l LoggerStub) Log(toLog entities.Log) (string, error) {
	if l.WithError {
		return "", errors.New("error in logger stub")
	}

	return fmt.Sprintf("Log handled for:%s", toLog.Name), nil
}
