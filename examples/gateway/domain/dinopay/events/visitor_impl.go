package events

import (
	"fmt"
	"mario/examples/gateway/domain"
	events2 "mario/examples/gateway/domain/events"
)

type VisitorImpl struct {
	gatewayEventsVisitor events2.Visitor
}

func (v VisitorImpl) VisitPaymentStatusUpdated(paymentStatusUpdated PaymentStatusUpdated) error {
	err := v.gatewayEventsVisitor.VisitDinopayPaymentUpdated(events2.DinopayPaymentUpdated{
		BaseEvent:              domain.BaseEvent{},
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
