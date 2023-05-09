package tests

import (
	"errors"
	"log-service/ports"
)

type LogRepositoryStub struct {
	WithError bool
}

func (l LogRepositoryStub) Insert(entry ports.LogEntry) error {
	if l.WithError {
		return errors.New("error in log repository stub")
	}

	return nil
}
