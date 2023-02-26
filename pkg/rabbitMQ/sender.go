package rabbitMQ

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type Sender struct {
	ch *amqp091.Channel
}

func NewSender(ch *amqp091.Channel) *Sender {
	return &Sender{ch: ch}
}

func (s *Sender) Send(message []byte, id uuid.UUID) {

	q, err := s.ch.QueueDeclare(
		"photos",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = s.ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %v\n", id)
}
