package main

import (
	"flag"
	"log"

	"github.com/streadway/amqp"
	"github.com/vdntruong/rabbitmq/util"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	exchangeType := flag.String("type", "default", "type exchange to demo")

	switch *exchangeType {
	case "default":
	default:
		log.Println("Unknown Exchange Type")
	}
}
