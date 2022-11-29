package events

import (
	"mario"
)

type PaymentStatusUpdated struct {
	mario.BaseEvent

	PaymentId string
	ClientId  string
	Status    string
	Timestamp int64
}

func (paymentStatusUpdated PaymentStatusUpdated) Accept(visitor Visitor) error {
	return visitor.VisitPaymentStatusUpdated(paymentStatusUpdated)
}
