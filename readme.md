# RABBIT MQ

### By default, RabbitMQ declare a nameless exchange ("") type called Default, routing my keys match with queue name.

```
  err = ch.Publish(
	"",     // exchange | nameless exchange
	q.Name, // routing key | queue's name
	false,  // mandatory
	false,  // immediate
	amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(MESSAGE),
	})
```
