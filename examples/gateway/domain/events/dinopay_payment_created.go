package events

import (
	"mario"
	"mario/examples/gateway/domain"
)

type DinopayPaymentCreated struct {
	domain.BaseEvent

	PaymentapiWithdrawalId string
	DinopayId              string
	DinopayStatus          string
	DinopayTime            int64
}

var _ mario.Event[Visitor] = DinopayPaymentCreated{}

func (dinopayPaymentCreated DinopayPaymentCreated) Accept(visitor Visitor) error {
	return visitor.VisitDinopayPaymentCreated(dinopayPaymentCreated)
}
