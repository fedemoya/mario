package events

import (
	"fmt"
	"mario"
	"mario/examples/gateway/domain/dinopay"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
)

type VisitorImpl struct {
	dinopayClient                dinopay.Client
	dinopayPaymentCreatedBuilder gatewayDomainEvents.DinopayPaymentCreatedBuilder
	cloudeventRepository         mario.CloudEventRepository
}

func NewVisitorImpl(
	dinopayClient dinopay.Client,
	dinopayPaymentCreatedBuilder gatewayDomainEvents.DinopayPaymentCreatedBuilder,
	cloudEventRepository mario.CloudEventRepository,
) *VisitorImpl {
	return &VisitorImpl{
		dinopayClient:                dinopayClient,
		dinopayPaymentCreatedBuilder: dinopayPaymentCreatedBuilder,
		cloudeventRepository:         cloudEventRepository,
	}
}

func (e VisitorImpl) VisitWithdrawalCreated(withdrawalCreated WithdrawalCreated) error {
	createPaymentRequest := dinopay.CreatePaymentRequest{
		SourceAccount:      withdrawalCreated.SourceAccount,
		DestinationAccount: withdrawalCreated.DestinationAccount,
		Amount:             withdrawalCreated.Amount,
		ClientID:           withdrawalCreated.Id,
	}

	res, err := e.dinopayClient.CreatePayment(createPaymentRequest)
	if err != nil {
		return fmt.Errorf("failed creating payment: %w", err)
	}

	dinopayPaymentCreated, err := e.dinopayPaymentCreatedBuilder.
		DinopayId(res.PaymentId).
		DinopayStatus(res.Status).DinopayTime(res.Time).
		PaymentapiWithdrawalId(withdrawalCreated.Id).
		Build()

	if err != nil {
		return fmt.Errorf("failed creating DinopayPaymentUpdated event: %w", err)
	}

	err = e.cloudeventRepository.Add(dinopayPaymentCreated)

	if err != nil {
		return fmt.Errorf("failed adding dinopayPaymentCreated event to the repository: %w", err)
	}

	return nil
}
