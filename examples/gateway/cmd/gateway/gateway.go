package main

import (
	"fmt"
	"mario"
	dinopayHttp "mario/examples/gateway/adapters/dinopay/http"
	"mario/examples/gateway/adapters/paymentapi/amqp"
	"mario/examples/gateway/adapters/paymentapi/events"
	paymentApiHttp "mario/examples/gateway/adapters/paymentapi/http"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
	paymentapiDomainEvents "mario/examples/gateway/domain/paymentapi/events"
	"time"
)

func main() {

	gatewayDomainEventsDispatchingVisitor := gatewayDomainEvents.NewDispatchingVisitor(paymentApiHttp.NewClient())

	paymentApiEventsVisitor := paymentapiDomainEvents.NewVisitorImpl(
		dinopayHttp.NewClient(),
		gatewayDomainEventsDispatchingVisitor,
	)

	paymentApiEventsSource := amqp.NewEventsSource()
	paymentApiEventsFactory := events.NewFactory()

	paymentApiEventsProcessor := mario.NewProcessor[paymentapiDomainEvents.Visitor](
		paymentApiEventsSource,
		paymentApiEventsFactory,
		paymentApiEventsVisitor,
		func(err error) {
			fmt.Printf("processor error: %s\n", err.Error())
		},
	)

	paymentApiEventsProcessor.Start()

	fmt.Println("Gateway started")

	time.Sleep(10 * time.Minute)
}
