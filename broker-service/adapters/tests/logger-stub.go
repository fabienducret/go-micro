package tests

import (
	"broker/ports"
	"fmt"
)

type loggerStub struct {
	Error error
}

func NewLoggerStub() *loggerStub {
	return &loggerStub{}
}

func (l *loggerStub) Log(toLog ports.Log) (string, error) {
	return fmt.Sprintf("Log handled for:%s", toLog.Name), l.Error
}
