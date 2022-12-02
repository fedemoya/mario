package events

import (
	"mario"
	"mario/examples/gateway/domain/dinopay"
	gatewayEvents "mario/examples/gateway/domain/gateway/events"
	"mario/examples/gateway/domain/paymentapi/events/handlers"
)

type VisitorImpl struct {
	withdrawalCreatedHandler *handlers.WithdrawalCreatedHandler
}

func NewVisitorImpl(
	dinopayClient dinopay.Client,
	dinopayEventsBuilderFactory gatewayEvents.BuildersFactory,
	cloudEventRepository mario.CloudEventRepository,
) *VisitorImpl {

	return &VisitorImpl{
		withdrawalCreatedHandler: handlers.NewWithdrawalCreatedHandler(
			dinopayClient,
			dinopayEventsBuilderFactory,
			cloudEventRepository,
		),
	}
}

func (e VisitorImpl) VisitWithdrawalCreated(withdrawalCreated WithdrawalCreated) error {
	return e.withdrawalCreatedHandler.Handle(withdrawalCreated)
}
