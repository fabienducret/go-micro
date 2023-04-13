package repositories

import (
	"broker/event"
	"broker/ports"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type loggerRepository struct {
	Rabbit *amqp.Connection
}

type logPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func NewLoggerRepository(rabbitConn *amqp.Connection) *loggerRepository {
	return &loggerRepository{
		Rabbit: rabbitConn,
	}
}

func (l *loggerRepository) Log(toLog ports.Log) error {
	emitter, err := event.NewEventEmitter(l.Rabbit)
	if err != nil {
		return err
	}

	payload := formatLogPayload(toLog)

	err = emitter.Push(payload, "log.INFO")
	if err != nil {
		return err
	}

	return nil
}

func formatLogPayload(toLog ports.Log) string {
	payload := logPayload(toLog)

	j, _ := json.MarshalIndent(&payload, "", "\t")

	return string(j)
}
