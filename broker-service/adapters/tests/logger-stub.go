package tests

import (
	"broker/ports"
	"errors"
	"fmt"
)

type LoggerStub struct {
	WithError bool
}

func (l LoggerStub) Log(toLog ports.Log) (string, error) {
	if l.WithError {
		return "", errors.New("error in logger stub")
	}

	return fmt.Sprintf("Log handled for:%s", toLog.Name), nil
}
