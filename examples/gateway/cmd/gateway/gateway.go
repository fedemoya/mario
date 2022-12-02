package main

import (
	"fmt"
	"mario"
	"mario/cloudevents/amqp/amqp"
	"mario/cloudevents/memdb"
	dinopayHttp "mario/examples/gateway/adapters/dinopay/http"
	events2 "mario/examples/gateway/adapters/gateway/events"
	paymentapiEvents "mario/examples/gateway/adapters/paymentapi/events"
	paymentapiHttp "mario/examples/gateway/adapters/paymentapi/http"
	"mario/examples/gateway/domain/gateway/events"
	paymentapiDomainEvents "mario/examples/gateway/domain/paymentapi/events"
	"time"
)

func main() {

	cloudEventBuilder := mario.NewCloudEventBuilderImpl()
	db := memdb.InitDB()
	cloudEventRepository := memdb.NewRepository(db, cloudEventBuilder)

	paymentapiEventsVisitor := paymentapiDomainEvents.NewVisitorImpl(
		dinopayHttp.NewClient(),
		events2.NewDinopayPaymentCreatedBuilder(cloudEventBuilder),
		cloudEventRepository,
	)

	paymentApiEventsSource := amqp.NewEventsSource()
	paymentApiEventsFactory := paymentapiEvents.NewFactory()

	paymentApiEventsProcessor := mario.NewProcessor[paymentapiDomainEvents.Visitor](
		paymentApiEventsSource,
		paymentApiEventsFactory,
		paymentapiEventsVisitor,
		func(err error) {
			fmt.Printf("paymentapi events processor error: %s\n", err.Error())
		},
	)

	paymentApiEventsProcessor.Start()

	gatewayDomainEventsVisitor := events.NewVisitorImpl(paymentapiHttp.NewClient())
	gatewayEventsReader := mario.NewCloudEventsReader(cloudEventRepository, events.GatewayCloudEventsSource)
	gatewayEventsFactory := events2.NewEventsFactory(memdb.Acknowledger{})

	gatewayEventsProcessor := mario.NewProcessor[events.Visitor](
		gatewayEventsReader,
		gatewayEventsFactory,
		gatewayDomainEventsVisitor,
		func(err error) {
			fmt.Printf("gateway events processor error: %s\n", err.Error())
		},
	)

	gatewayEventsProcessor.Start()

	fmt.Println("Gateway started")

	time.Sleep(10 * time.Minute)
}
