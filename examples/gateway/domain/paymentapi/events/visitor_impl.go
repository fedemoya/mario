package events

import (
	"fmt"
	"mario/examples/gateway/domain"
	"mario/examples/gateway/domain/dinopay"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
)

type VisitorImpl struct {
	dinopayClient        dinopay.Client
	gatewayEventsVisitor gatewayDomainEvents.Visitor
}

func NewVisitorImpl(dinopayClient dinopay.Client, gatewayEventsVisitor gatewayDomainEvents.Visitor) *VisitorImpl {
	return &VisitorImpl{
		dinopayClient:        dinopayClient,
		gatewayEventsVisitor: gatewayEventsVisitor,
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

	err = e.gatewayEventsVisitor.VisitDinopayPaymentCreated(gatewayDomainEvents.DinopayPaymentCreated{
		BaseEvent:              domain.BaseEvent{},
		PaymentapiWithdrawalId: withdrawalCreated.Id,
		DinopayId:              res.PaymentId,
		DinopayStatus:          res.Status,
		DinopayTime:            res.Time,
	})
	if err != nil {
		return fmt.Errorf("failed visiting DinopayPaymentUpdated event: %w", err)
	}

	return nil
}
