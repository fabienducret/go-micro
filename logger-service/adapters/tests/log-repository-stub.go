package tests

import (
	"errors"
	"log-service/entities"
)

type LogRepositoryStub struct {
	WithError bool
}

func (l LogRepositoryStub) Insert(entry entities.LogEntry) error {
	if l.WithError {
		return errors.New("error in log repository stub")
	}

	return nil
}
