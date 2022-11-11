package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"mario"
)

// TODO improve

type EventsSource struct {
	subscriptions []chan mario.RawEvent
}

func NewEventsSource() *EventsSource {
	return &EventsSource{subscriptions: make([]chan mario.RawEvent, 0)}
}

func (es *EventsSource) Start() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

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
				ch <- d.Body
			}
		}
	}()

	return nil
}

func (es *EventsSource) Subscribe() (<-chan mario.RawEvent, <-chan error) {
	ch := make(chan mario.RawEvent)
	es.subscriptions = append(es.subscriptions, ch)
	return ch, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
