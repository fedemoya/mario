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
	PaymentapiWithdrawalId(paymentapiWithdrawalId string) DinopayPaymentCreatedBuilder
	DinopayId(dinopayId string) DinopayPaymentCreatedBuilder
	DinopayStatus(dinopayStatus string) DinopayPaymentCreatedBuilder
	DinopayTime(dinopayTime int64) DinopayPaymentCreatedBuilder
	CorrelationID(id string) DinopayPaymentCreatedBuilder
	Build() (DinopayPaymentCreated, error)
}
