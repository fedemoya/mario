package events

import (
	"mario"
	gatewayEvents "mario/examples/gateway/domain/gateway/events"
)

type BuildersFactory struct {
	cloudEventBuilder mario.CloudEventBuilder
	acknowledger      mario.Acknowledger
}

func NewBuildersFactory(
	cloudEventBuilder mario.CloudEventBuilder,
	acknowledger mario.Acknowledger,
) *BuildersFactory {

	return &BuildersFactory{
		cloudEventBuilder: cloudEventBuilder,
		acknowledger:      acknowledger,
	}
}

func (bf *BuildersFactory) CreateDinopayPaymentCreatedBuilder() gatewayEvents.DinopayPaymentCreatedBuilder {
	return NewDinopayPaymentCreatedBuilder(
		bf.cloudEventBuilder,
		bf.acknowledger,
	)
}

func (bf *BuildersFactory) CreateDinopayPaymentUpdatedBuilder() gatewayEvents.DinopayPaymentUpdatedBuilder {
	return NewDinopayPaymentUpdatedBuilder(
		bf.cloudEventBuilder,
		bf.acknowledger,
	)
}
