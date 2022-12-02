package events

import (
	"fmt"
	"mario"
	"mario/examples/gateway/domain/gateway/events"
)

type VisitorImpl struct {
	gatewayEventsVisitor events.Visitor
}

func (v VisitorImpl) VisitPaymentStatusUpdated(paymentStatusUpdated PaymentStatusUpdated) error {
	err := v.gatewayEventsVisitor.VisitDinopayPaymentUpdated(events.DinopayPaymentUpdated{
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
