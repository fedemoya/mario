package amqp

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"mario"
	cloudEventsJson "mario/cloudevents/json"
)

// TODO improve

type EventsReader struct {
	subscriptions []chan mario.CloudEvent
}

func NewEventsReader() *EventsReader {
	return &EventsReader{subscriptions: make([]chan mario.CloudEvent, 0)}
}

func (es *EventsReader) Start() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	// TODO close connection

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// TODO close channel

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			for _, ch := range es.subscriptions {
				var cloudEvent cloudEventsJson.CloudEvent
				err := json.Unmarshal(d.Body, &cloudEvent)
				// TODO send err over err channel
				if err != nil {
					panic(err)
				}
				ch <- cloudEvent
			}
		}
	}()

	return nil
}

func (es *EventsReader) Subscribe() (<-chan mario.CloudEvent, <-chan error) {
	ch := make(chan mario.CloudEvent)
	es.subscriptions = append(es.subscriptions, ch)
	return ch, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}