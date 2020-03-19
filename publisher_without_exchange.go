package main

import (
	"os"

	"github.com/streadway/amqp"
	"github.com/vdntruong/rabbitmq/util"
)

// Publisher, publish to queue without exchange
func DefaultPublisher() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// a channel, which is where most of the API for getting things done resides
	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("queue", false, false, false, false, nil)
	util.FailOnError(err, "Failed to declare a queue")

	body := util.BodyFrom(os.Args)
	err = ch.Publish(
		"",
		q.Name, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	util.FailOnError(err, "Failed to publish message")
}
