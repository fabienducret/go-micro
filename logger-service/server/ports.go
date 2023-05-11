package server

import "log-service/entities"

type LogRepository interface {
	Insert(entities.LogEntry) error
}
