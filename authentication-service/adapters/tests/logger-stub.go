package tests

import "authentication/ports"

type loggerStub struct{}

func NewLoggerStub() *loggerStub {
	return &loggerStub{}
}

func (r *loggerStub) Log(toLog ports.Log) error {
	return nil
}
