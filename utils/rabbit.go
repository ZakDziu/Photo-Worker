package utils

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"main/pkg/rabbitMQ"
)

type Bus struct {
	Receiver *rabbitMQ.Receiver
	Sender   *rabbitMQ.Sender
	Conn     *amqp.Connection
	Ch       *amqp.Channel
}

func StartMQWorker() *Bus {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	rabbitMQ.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	rabbitMQ.FailOnError(err, "Failed to open a channel")

	rec := rabbitMQ.NewReceiver(ch)
	sen := rabbitMQ.NewSender(ch)
	return &Bus{
		Receiver: rec,
		Sender:   sen,
		Conn:     conn,
		Ch:       ch,
	}
}
