package repositories

import (
	"broker/ports"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type loggerRepository struct{}

type logPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

const logServiceUrl = "http://logger-service/log"

func NewLoggerRepository() *loggerRepository {
	return &loggerRepository{}
}

func (l *loggerRepository) Log(toLog ports.Log) error {
	toSend := formatLogRequest(toLog)

	request, err := http.NewRequest("POST", logServiceUrl, toSend)
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

func formatLogRequest(toLog ports.Log) *bytes.Buffer {
	entry := logPayload{
		Name: toLog.Name,
		Data: toLog.Data,
	}

	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	return bytes.NewBuffer(jsonData)
}
