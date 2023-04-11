package repositories

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const logServiceURL = "http://logger-service/log"

type loggerRepository struct{}

func NewLoggerRepository() *loggerRepository {
	return &loggerRepository{}
}

func (r *loggerRepository) Log(name, data string) error {
	toSend := formatLoggerRequest(name, data)

	request, err := http.NewRequest("POST", logServiceURL, toSend)
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

func formatLoggerRequest(name, data string) *bytes.Buffer {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	return bytes.NewBuffer(jsonData)
}
