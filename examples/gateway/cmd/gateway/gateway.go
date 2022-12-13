package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"mario"
	"mario/cloudevents/amqp"
	"mario/cloudevents/memdb"
	dinopayHttp "mario/examples/gateway/adapters/dinopay/http"
	gatewayEventsAdapters "mario/examples/gateway/adapters/gateway/events"
	paymentapiEvents "mario/examples/gateway/adapters/paymentapi/events"
	paymentapiHttp "mario/examples/gateway/adapters/paymentapi/http"
	gatewayDomainEvents "mario/examples/gateway/domain/gateway/events"
	paymentapiDomainEvents "mario/examples/gateway/domain/paymentapi/events"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger := zerolog.Logger{}

	cloudEventBuilder := mario.NewCloudEventBuilderImpl()

	db := memdb.InitDB()
	cloudEventRepository := memdb.NewRepository(db, cloudEventBuilder)

	paymentApiEventsProcessor := buildPaymentapiEventsProcessor(cloudEventBuilder, cloudEventRepository)
	err := paymentApiEventsProcessor.Start()
	if err != nil {
		panic(err)
	}

	gatewayEventsReaderCtx := context.Background()
	gatewayEventsProcessor := buildGatewayEventsProcessor(gatewayEventsReaderCtx, cloudEventRepository)
	err = gatewayEventsProcessor.Start()
	if err != nil {
		panic(err)
	}

	fmt.Println("Gateway started")

	signalsChannel := make(chan os.Signal, 1)
	signal.Notify(signalsChannel, syscall.SIGINT, syscall.SIGTERM)
	s := <-signalsChannel
	logger.Info().Msgf("Received shutdown signal %s. Shutting down...", s)

	// TODO make all implement a Service interface with a Start(ctx) method
	// TODO and stop them in a for loop
	paymentApiEventsProcessor.Stop()
	gatewayEventsReaderCtx.Done()
	gatewayEventsProcessor.Stop()
}

func buildPaymentapiEventsProcessor(cloudEventBuilder *mario.CloudEventBuilderImpl, cloudEventRepository *memdb.Repository) *mario.Processor[paymentapiDomainEvents.Visitor] {
	repositoryAcknowledger := mario.NewRepositoryAcknowledger(cloudEventRepository, 5)

	paymentapiEventsVisitor := paymentapiDomainEvents.NewVisitorImpl(
		dinopayHttp.NewClient(),
		gatewayEventsAdapters.NewBuildersFactory(cloudEventBuilder, repositoryAcknowledger),
		cloudEventRepository,
	)

	paymentApiEventsReader := amqp.NewEventsReader()
	paymentApiEventsReader.Start()

	paymentApiEventsFactory := paymentapiEvents.NewFactory()

	paymentApiEventsProcessor := mario.NewProcessor[paymentapiDomainEvents.Visitor](
		paymentApiEventsReader,
		paymentApiEventsFactory,
		paymentapiEventsVisitor,
		func(err error) {
			fmt.Printf("paymentapi events processor error: %s\n", err.Error())
		},
	)
	return paymentApiEventsProcessor
}

func buildGatewayEventsProcessor(gatewayEventsReaderCtx context.Context, cloudEventRepository *memdb.Repository) *mario.Processor[gatewayDomainEvents.Visitor] {
	repositoryAcknowledger := mario.NewRepositoryAcknowledger(cloudEventRepository, 5)

	gatewayEventsFactory := gatewayEventsAdapters.NewEventsFactory(repositoryAcknowledger)
	gatewayDomainEventsVisitor := gatewayDomainEvents.NewVisitorImpl(paymentapiHttp.NewClient())

	gatewayEventsReader := mario.NewCloudEventsReader(cloudEventRepository, gatewayDomainEvents.GatewayCloudEventsSource)
	gatewayEventsReader.Start(gatewayEventsReaderCtx)

	gatewayEventsProcessor := mario.NewProcessor[gatewayDomainEvents.Visitor](
		gatewayEventsReader,
		gatewayEventsFactory,
		gatewayDomainEventsVisitor,
		func(err error) {
			fmt.Printf("gateway events processor error: %s\n", err.Error())
		},
	)

	return gatewayEventsProcessor
}
