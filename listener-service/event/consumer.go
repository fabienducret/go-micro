package event

import (
	"encoding/json"
	"fmt"
	"listener/ports"
	"listener/repositories"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

func (consumer *Consumer) Listen(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, s := range topics {
		err = ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil,
		)

		if err != nil {
			return nil
		}
	}

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil
	}

	forever := make(chan bool)
	go func() {
		lr := repositories.NewLoggerRepository()
		for d := range messages {
			var payload ports.Payload
			_ = json.Unmarshal(d.Body, &payload)
			log.Println("New message to handle", payload.Name)

			go handlePayload(lr, payload)
		}
	}()

	fmt.Printf("Waiting for message on [Exchange, Queue] [logs_topic, %s]\n", q.Name)
	<-forever

	return nil
}

func handlePayload(lr ports.Logger, payload ports.Payload) {
	switch payload.Name {
	case "log", "event":
		toLog := ports.Log(payload)
		err := lr.Log(toLog)
		if err != nil {
			log.Println(err)
		}

	case "auth":
	default:
		toLog := ports.Log(payload)
		err := lr.Log(toLog)
		if err != nil {
			log.Println(err)
		}
	}
}
