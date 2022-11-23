package dinopay_payment_created

import (
	"fmt"
	"github.com/google/uuid"
	"mario"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
	"time"
)

type DinopayPaymentCreatedBuilder struct {
	cloudEventBuilder mario.CloudEventBuilder

	paymentapiWithdrawalId string
	dinopayId              string
	dinopayStatus          string
	dinopayTime            int64
	correlationID          string
	acknowledger           mario.Acknowledger
}

func NewDinopayPaymentCreatedBuilder(cloudEventBuilder mario.CloudEventBuilder) *DinopayPaymentCreatedBuilder {
	return &DinopayPaymentCreatedBuilder{
		cloudEventBuilder: cloudEventBuilder,
	}
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

func (b *DinopayPaymentCreatedBuilder) Acknowledger(acknowledger mario.Acknowledger) gatewayDomainEvents.DinopayPaymentCreatedBuilder {
	b.acknowledger = acknowledger
	return b
}

func (b *DinopayPaymentCreatedBuilder) Build() (gatewayDomainEvents.DinopayPaymentCreated, error) {
	event := gatewayDomainEvents.DinopayPaymentCreated{
		PaymentapiWithdrawalId: b.paymentapiWithdrawalId,
		DinopayId:              b.dinopayId,
		DinopayStatus:          b.dinopayStatus,
		DinopayTime:            b.dinopayTime,
	}
	eventJson, err := marshalJSON(event)
	if err != nil {
		return gatewayDomainEvents.DinopayPaymentCreated{}, fmt.Errorf("failed building DinopayPaymentCreated event: %w", err)
	}
	cloudEvent, _ := b.cloudEventBuilder.
		Id(uuid.New().String()).
		Source(gatewayDomainEvents.GatewayCloudEventsSource).
		EventType(gatewayDomainEvents.EventTypeDinopayPaymentCreated).
		CorrelationID(b.correlationID).
		Time(time.Now().Unix()).
		Data(eventJson).
		Build()
	baseEvent := mario.NewBaseEvent(
		cloudEvent,
		b.acknowledger,
	)
	event.BaseEvent = baseEvent
	return event, nil
}
