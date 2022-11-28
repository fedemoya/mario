package events

import (
	"mario"
)

type DinopayPaymentUpdated struct {
	mario.BaseEvent

	PaymentapiWithdrawalId string
	DinopayId              string
	DinopayStatus          string
	DinopayTime            int64
}

var _ mario.Event[Visitor] = DinopayPaymentUpdated{}

func (dinopayPaymentUpdated DinopayPaymentUpdated) Accept(visitor Visitor) error {
	return visitor.VisitDinopayPaymentUpdated(dinopayPaymentUpdated)
}
