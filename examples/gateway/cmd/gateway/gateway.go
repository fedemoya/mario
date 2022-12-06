package main

import (
	"fmt"
	"mario"
	"mario/cloudevents/amqp/amqp"
	"mario/cloudevents/memdb"
	dinopayHttp "mario/examples/gateway/adapters/dinopay/http"
	gatewayEventsAdapters "mario/examples/gateway/adapters/gateway/events"
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
	repositoryAcknowledger := mario.NewRepositoryAcknowledger(cloudEventRepository, 5)

	paymentapiEventsVisitor := paymentapiDomainEvents.NewVisitorImpl(
		dinopayHttp.NewClient(),
		gatewayEventsAdapters.NewBuildersFactory(cloudEventBuilder, repositoryAcknowledger),
		cloudEventRepository,
	)

	paymentApiEventsReader := amqp.NewEventsReader()
	paymentApiEventsFactory := paymentapiEvents.NewFactory()

	paymentApiEventsProcessor := mario.NewProcessor[paymentapiDomainEvents.Visitor](
		paymentApiEventsReader,
		paymentApiEventsFactory,
		paymentapiEventsVisitor,
		func(err error) {
			fmt.Printf("paymentapi events processor error: %s\n", err.Error())
		},
	)

	paymentApiEventsProcessor.Start()

	gatewayDomainEventsVisitor := events.NewVisitorImpl(paymentapiHttp.NewClient())
	gatewayEventsReader := mario.NewCloudEventsReader(cloudEventRepository, events.GatewayCloudEventsSource)
	gatewayEventsFactory := gatewayEventsAdapters.NewEventsFactory(repositoryAcknowledger)

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
