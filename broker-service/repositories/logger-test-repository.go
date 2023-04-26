package repositories

import (
	"broker/ports"
	"fmt"
)

type loggerTestRepository struct{}

func NewLoggerTestRepository() *loggerTestRepository {
	return &loggerTestRepository{}
}

func (l *loggerTestRepository) Log(toLog ports.Log) (string, error) {
	return fmt.Sprintf("Log handled for:%s", toLog.Name), nil
}
