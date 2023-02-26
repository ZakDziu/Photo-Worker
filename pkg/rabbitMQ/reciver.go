package rabbitMQ

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"main/internal/models"
	"main/internal/services"
	"sync"
)

type Receiver struct {
	ch *amqp091.Channel
}

func NewReceiver(ch *amqp091.Channel) *Receiver {
	return &Receiver{ch: ch}
}

func (r *Receiver) Listen(group *sync.WaitGroup) {
	defer group.Done()
	q, err := r.ch.QueueDeclare(
		"photos",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := r.ch.Consume(
		q.Name, // queue
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var photo models.Photo
			if err := json.Unmarshal(d.Body, &photo); err != nil {
				FailOnError(err, "Failed to unmarshal body")
			}

			services.OptimizePhoto(photo.ID)
			log.Printf(" [*] Received a file with id %v", photo.ID)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
