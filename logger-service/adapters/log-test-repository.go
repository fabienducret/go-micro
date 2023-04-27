package adapters

import (
	"log-service/ports"
)

type logTestRepository struct {
}

func NewLogTestRepository() *logTestRepository {
	return &logTestRepository{}
}

func (r *logTestRepository) Insert(entry ports.LogEntry) error {
	return nil
}
