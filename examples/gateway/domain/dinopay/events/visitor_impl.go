package events

import (
	"fmt"
	"mario"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
)

type VisitorImpl struct {
	gatewayEventsVisitor gatewayDomainEvents.Visitor
}

func (v VisitorImpl) VisitPaymentStatusUpdated(paymentStatusUpdated PaymentStatusUpdated) error {
	err := v.gatewayEventsVisitor.VisitDinopayPaymentUpdated(gatewayDomainEvents.DinopayPaymentUpdated{
		BaseEvent:              mario.BaseEvent{},
		PaymentapiWithdrawalId: paymentStatusUpdated.ClientId,
		DinopayId:              paymentStatusUpdated.PaymentId,
		DinopayStatus:          paymentStatusUpdated.Status,
		DinopayTime:            paymentStatusUpdated.Timestamp,
	})
	if err != nil {
		return fmt.Errorf("failed visiting event DinopayPaymentUpdated: %w", err)
	}

	return nil
}
