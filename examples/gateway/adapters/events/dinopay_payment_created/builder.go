package dinopay_payment_created

import (
	"github.com/google/uuid"
	"mario"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
	"time"
)

type DinopayPaymentCreatedBuilder struct {
	paymentapiWithdrawalId string
	dinopayId              string
	dinopayStatus          string
	dinopayTime            int64
	correlationID          string
}

func NewDinopayPaymentCreatedBuilder() *DinopayPaymentCreatedBuilder {
	return &DinopayPaymentCreatedBuilder{}
}

func (b *DinopayPaymentCreatedBuilder) PaymentapiWithdrawalId(paymentapiWithdrawalId string) gatewayDomainEvents.DinopayPaymentCreatedBuilder {
	b.paymentapiWithdrawalId = paymentapiWithdrawalId
	return b
}

func (b *DinopayPaymentCreatedBuilder) DinopayId(dinopayId string) gatewayDomainEvents.DinopayPaymentCreatedBuilder {
	b.dinopayId = dinopayId
	return b
}

func (b *DinopayPaymentCreatedBuilder) DinopayStatus(dinopayStatus string) gatewayDomainEvents.DinopayPaymentCreatedBuilder {
	b.dinopayStatus = dinopayStatus
	return b
}

func (b *DinopayPaymentCreatedBuilder) DinopayTime(dinopayTime int64) gatewayDomainEvents.DinopayPaymentCreatedBuilder {
	b.dinopayTime = dinopayTime
	return b
}

func (b *DinopayPaymentCreatedBuilder) CorrelationID(correlationID string) gatewayDomainEvents.DinopayPaymentCreatedBuilder {
	b.correlationID = correlationID
	return b
}

func (b *DinopayPaymentCreatedBuilder) Build() (gatewayDomainEvents.DinopayPaymentCreated, error) {
	event := &gatewayDomainEvents.DinopayPaymentCreated{
		PaymentapiWithdrawalId: b.paymentapiWithdrawalId,
		DinopayId:              b.dinopayId,
		DinopayStatus:          b.dinopayStatus,
		DinopayTime:            b.dinopayTime,
	}
	baseEvent := mario.NewBaseEvent(
		cloudEvent{
			id:             uuid.New().String(),
			source:         gatewayDomainEvents.GatewayCloudEventsSource,
			cloudEventType: gatewayDomainEvents.EventTypeDinopayPaymentCreated,
			time:           time.Now().Unix(),
			correlationID:  b.correlationID,
		},
		mario.DummyAcknowledger{},
		jsonMarshaler{
			event: *event,
		},
	)
	event.BaseEvent = baseEvent
	return *event, nil
}
