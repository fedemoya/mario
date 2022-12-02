package handlers

import (
	"fmt"
	"mario"
	"mario/examples/gateway/domain/dinopay"
	gatewayEvents "mario/examples/gateway/domain/gateway/events"
	"mario/examples/gateway/domain/paymentapi/events"
)

type WithdrawalCreatedHandler struct {
	dinopayClient               dinopay.Client
	dinopayEventsBuilderFactory gatewayEvents.BuildersFactory
	cloudeventRepository        mario.CloudEventRepository
}

func NewWithdrawalCreatedHandler(
	dinopayClient dinopay.Client,
	dinopayEventsBuilderFactory gatewayEvents.BuildersFactory,
	cloudEventRepository mario.CloudEventRepository,
) *WithdrawalCreatedHandler {

	return &WithdrawalCreatedHandler{
		dinopayClient:               dinopayClient,
		dinopayEventsBuilderFactory: dinopayEventsBuilderFactory,
		cloudeventRepository:        cloudEventRepository,
	}
}

func (w *WithdrawalCreatedHandler) Handle(withdrawalCreated events.WithdrawalCreated) error {
	createPaymentRequest := dinopay.CreatePaymentRequest{
		SourceAccount:      withdrawalCreated.SourceAccount,
		DestinationAccount: withdrawalCreated.DestinationAccount,
		Amount:             withdrawalCreated.Amount,
		ClientID:           withdrawalCreated.Id,
	}

	res, err := w.dinopayClient.CreatePayment(createPaymentRequest)
	if err != nil {
		return fmt.Errorf("failed creating payment: %w", err)
	}

	dinopayPaymentCreated, err := w.dinopayEventsBuilderFactory.
		CreateDinopayPaymentCreatedBuilder().
		DinopayId(res.PaymentId).
		DinopayStatus(res.Status).DinopayTime(res.Time).
		PaymentapiWithdrawalId(withdrawalCreated.Id).
		Build()

	if err != nil {
		return fmt.Errorf("failed creating DinopayPaymentUpdated event: %w", err)
	}

	err = w.cloudeventRepository.Add(dinopayPaymentCreated)

	if err != nil {
		return fmt.Errorf("failed adding dinopayPaymentCreated event to the repository: %w", err)
	}

	return nil
}