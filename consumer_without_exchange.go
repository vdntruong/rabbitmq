package main

import (
	"log"

	"github.com/streadway/amqp"
	"github.com/vdntruong/rabbitmq/util"
)

func DefaultConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("queue", false, false, false, false, nil)
	util.FailOnError(err, "Failed to declare a queue")

	q2, err := ch.QueueDeclare("queue", false, false, false, false, nil)
	util.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "consumer01", true, false, false, false, nil)
	util.FailOnError(err, "Failed to register a consumer")

	msgs2, err := ch.Consume(q.Name, "consumer02", true, false, false, false, nil)
	util.FailOnError(err, "Failed to register a consumer")

	msgs3, err := ch.Consume(q2.Name, "consumer03", true, false, false, false, nil)
	util.FailOnError(err, "Failed to register a consumer")

	msgs4, err := ch.Consume(q2.Name, "consumer04", true, false, false, false, nil)
	util.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message 1: %s", d.Body)
		}
	}()

	go func() {
		for d := range msgs2 {
			log.Printf("Received a message 2: %s", d.Body)
		}
	}()

	go func() {
		for d := range msgs3 {
			log.Printf("Received a message 3: %s", d.Body)
		}
	}()

	go func() {
		for d := range msgs4 {
			log.Printf("Received a message 4: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
