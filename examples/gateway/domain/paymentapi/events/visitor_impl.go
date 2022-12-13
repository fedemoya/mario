package events

import (
	"github.com/rs/zerolog/log"
	"mario"
	"mario/examples/gateway/domain/dinopay"
	gatewayEvents "mario/examples/gateway/domain/gateway/events"
)

type VisitorImpl struct {
	withdrawalCreatedHandler *WithdrawalCreatedHandler
}

func NewVisitorImpl(
	dinopayClient dinopay.Client,
	dinopayEventsBuilderFactory gatewayEvents.BuildersFactory,
	cloudEventRepository mario.CloudEventRepository,
) *VisitorImpl {

	return &VisitorImpl{
		withdrawalCreatedHandler: NewWithdrawalCreatedHandler(
			dinopayClient,
			dinopayEventsBuilderFactory,
			cloudEventRepository,
		),
	}
}

func (e VisitorImpl) VisitWithdrawalCreated(withdrawalCreated WithdrawalCreated) error {
	log.Info().Msgf("visiting %s event with Id %s", withdrawalCreated.Type(), withdrawalCreated.Id)
	return e.withdrawalCreatedHandler.Handle(withdrawalCreated)
}
