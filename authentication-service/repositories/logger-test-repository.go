package repositories

import "authentication/ports"

type loggerTestRepository struct{}

func NewLoggerTestRepository() *loggerTestRepository {
	return &loggerTestRepository{}
}

func (r *loggerTestRepository) Log(toLog ports.Log) error {
	return nil
}
