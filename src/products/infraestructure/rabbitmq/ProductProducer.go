package rabbitmq

import (
	"context"
	"log"
	"time"

	"mi-tienda-online/src/core"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishProduct(body string) {
	conn := core.GetRabbitMQConnection()
	defer conn.Close()

	ch, err := conn.Channel()
	core.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"HolaMundo", // name of the queue
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	core.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	core.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
