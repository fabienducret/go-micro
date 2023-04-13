package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"listener/ports"
	"net/http"
)

type loggerRepository struct{}

const logServiceUrl = "http://logger-service/log"

func NewLoggerRepository() *loggerRepository {
	return &loggerRepository{}
}

func (r *loggerRepository) Log(entry ports.Log) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	request, err := http.NewRequest("POST", logServiceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json ")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return errors.New("not accepted from logger service")
	}

	return nil
}
