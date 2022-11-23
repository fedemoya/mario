package events

import (
	"mario"
)

type DinopayPaymentCreated struct {
	mario.BaseEvent

	PaymentapiWithdrawalId string
	DinopayId              string
	DinopayStatus          string
	DinopayTime            int64
}

var _ mario.Event[Visitor] = DinopayPaymentCreated{}

func (dinopayPaymentCreated DinopayPaymentCreated) Accept(visitor Visitor) error {
	return visitor.VisitDinopayPaymentCreated(dinopayPaymentCreated)
}

type DinopayPaymentCreatedBuilder interface {
	PaymentapiWithdrawalId(paymentapiWithdrawalId string) DinopayPaymentCreatedBuilder
	DinopayId(dinopayId string) DinopayPaymentCreatedBuilder
	DinopayStatus(dinopayStatus string) DinopayPaymentCreatedBuilder
	DinopayTime(dinopayTime int64) DinopayPaymentCreatedBuilder
	Acknowledger(acknowledger mario.Acknowledger) DinopayPaymentCreatedBuilder
	Build() (DinopayPaymentCreated, error)
}
