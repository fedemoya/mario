package events

import (
	"mario/examples/gateway/domain"
)

type PaymentStatusUpdated struct {
	domain.BaseEvent

	PaymentId string
	ClientId  string
	Status    string
	Timestamp int64
}

func (paymentStatusUpdated PaymentStatusUpdated) Accept(visitor Visitor) error {
	return visitor.VisitPaymentStatusUpdated(paymentStatusUpdated)
}
