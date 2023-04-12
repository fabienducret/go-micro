package repositories

import (
	"broker/ports"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type mailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type mailerRepository struct{}

const mailServiceUrl = "http://mail-service/send"

func NewMailerRepository() *mailerRepository {
	return &mailerRepository{}
}

func (r *mailerRepository) Send(mail ports.Mail) error {
	toSend := formatMailRequest(mail)

	request, err := http.NewRequest("POST", mailServiceUrl, toSend)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return errors.New("error calling mail service")
	}

	return nil
}

func formatMailRequest(mail ports.Mail) *bytes.Buffer {
	toSend := mailPayload{
		From:    mail.From,
		To:      mail.To,
		Subject: mail.Subject,
		Message: mail.Message,
	}

	jsonData, _ := json.MarshalIndent(toSend, "", "\t")

	return bytes.NewBuffer(jsonData)
}
