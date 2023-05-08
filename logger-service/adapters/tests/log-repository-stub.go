package tests

import (
	"log-service/ports"
)

type logRepositoryStub struct {
}

func NewLogRepositoryStub() *logRepositoryStub {
	return &logRepositoryStub{}
}

func (r *logRepositoryStub) Insert(entry ports.LogEntry) error {
	return nil
}
