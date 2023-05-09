package tests

import (
	"broker/ports"
	"fmt"
)

type LoggerStub struct {
	Error error
}

func (l LoggerStub) Log(toLog ports.Log) (string, error) {
	return fmt.Sprintf("Log handled for:%s", toLog.Name), l.Error
}
