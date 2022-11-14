package main

import (
	"fmt"
	"mario"
	dinopayHttp "mario/examples/gateway/adapters/dinopay/http"
	"mario/examples/gateway/adapters/paymentapi/amqp"
	"mario/examples/gateway/adapters/paymentapi/events"
	paymentapiHttp "mario/examples/gateway/adapters/paymentapi/http"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
	paymentapiDomainEvents "mario/examples/gateway/domain/paymentapi/events"
	"time"
)

func main() {

	gatewayDomainEventsVisitor := gatewayDomainEvents.NewVisitorImpl(paymentapiHttp.NewClient())

	paymentapiEventsVisitor := paymentapiDomainEvents.NewVisitorImpl(
		dinopayHttp.NewClient(),
		gatewayDomainEventsVisitor,
	)

	paymentApiEventsSource := amqp.NewEventsSource()
	paymentApiEventsFactory := events.NewFactory()

	paymentApiEventsProcessor := mario.NewProcessor[paymentapiDomainEvents.Visitor](
		paymentApiEventsSource,
		paymentApiEventsFactory,
		paymentapiEventsVisitor,
		func(err error) {
			fmt.Printf("paymentapi events processor error: %s\n", err.Error())
		},
	)

	paymentApiEventsProcessor.Start()

	fmt.Println("Gateway started")

	time.Sleep(10 * time.Minute)
}
