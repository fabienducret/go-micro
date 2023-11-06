package logger

import (
	"log"
	"log-service/entities"
)

type Payload struct {
	Name string
	Data string
}

type LogRepository interface {
	Insert(entities.LogEntry) error
}

type Logger struct {
	LogRepository LogRepository
}

func New(lr LogRepository) *Logger {
	l := new(Logger)
	l.LogRepository = lr

	return l
}

func (l *Logger) LogInfo(payload Payload, resp *string) error {
	err := l.LogRepository.Insert(entities.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Log handled for:" + payload.Name

	return nil
}
