package main

import (
	"fmt"
	"mario"
	dinopayHttp "mario/examples/gateway/adapters/dinopay/http"
	"mario/examples/gateway/adapters/events/dinopay_payment_created"
	"mario/examples/gateway/adapters/paymentapi/amqp"
	paymentapiEvents "mario/examples/gateway/adapters/paymentapi/events"
	paymentapiHttp "mario/examples/gateway/adapters/paymentapi/http"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
	paymentapiDomainEvents "mario/examples/gateway/domain/paymentapi/events"
	"mario/memdb"
	"time"
)

func main() {

	cloudEventBuilder := mario.NewCloudEventBuilderImpl()
	db := memdb.InitDB()
	cloudEventRepository := memdb.NewRepository(db, cloudEventBuilder)

	paymentapiEventsVisitor := paymentapiDomainEvents.NewVisitorImpl(
		dinopayHttp.NewClient(),
		dinopay_payment_created.NewDinopayPaymentCreatedBuilder(cloudEventBuilder),
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

	gatewayDomainEventsVisitor := gatewayDomainEvents.NewVisitorImpl(paymentapiHttp.NewClient())
	gatewayEventsReader := mario.NewCloudEventsReader(cloudEventRepository, gatewayDomainEvents.GatewayCloudEventsSource)
	gatewayEventsFactory := dinopay_payment_created.NewEventsFactory(memdb.Acknowledger{})

	gatewayEventsProcessor := mario.NewProcessor[gatewayDomainEvents.Visitor](
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
