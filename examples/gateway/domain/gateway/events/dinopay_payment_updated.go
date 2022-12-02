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

type DinopayPaymentUpdatedBuilder interface {
	PaymentapiWithdrawalId(paymentapiWithdrawalId string) DinopayPaymentUpdatedBuilder
	DinopayId(dinopayId string) DinopayPaymentUpdatedBuilder
	DinopayStatus(dinopayStatus string) DinopayPaymentUpdatedBuilder
	DinopayTime(dinopayTime int64) DinopayPaymentUpdatedBuilder
	CorrelationID(id string) DinopayPaymentUpdatedBuilder
	Build() (DinopayPaymentUpdated, error)
}
