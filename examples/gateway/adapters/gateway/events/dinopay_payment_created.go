package events

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"mario"
	"mario/examples/gateway/domain/gateway/events"
	"time"
)

type dinopayPaymentCreated struct {
	PaymentapiWithdrawalId string `json:"paymentapi_withdrawal_id"`
	DinopayId              string `json:"dinopay_id"`
	DinopayStatus          string `json:"dinopay_status"`
	DinopayTime            int64  `json:"dinopay_time"`
}

type DinopayPaymentCreatedBuilder struct {
	cloudEventBuilder mario.CloudEventBuilder
	acknowledger      mario.Acknowledger

	paymentapiWithdrawalId string
	dinopayId              string
	dinopayStatus          string
	dinopayTime            int64
	correlationID          string
}

func NewDinopayPaymentCreatedBuilder(
	cloudEventBuilder mario.CloudEventBuilder,
	acknowledger mario.Acknowledger,
) *DinopayPaymentCreatedBuilder {

	return &DinopayPaymentCreatedBuilder{
		cloudEventBuilder: cloudEventBuilder,
		acknowledger:      acknowledger,
	}
}

func (b *DinopayPaymentCreatedBuilder) PaymentapiWithdrawalId(paymentapiWithdrawalId string) events.DinopayPaymentCreatedBuilder {
	b.paymentapiWithdrawalId = paymentapiWithdrawalId
	return b
}

func (b *DinopayPaymentCreatedBuilder) DinopayId(dinopayId string) events.DinopayPaymentCreatedBuilder {
	b.dinopayId = dinopayId
	return b
}

func (b *DinopayPaymentCreatedBuilder) DinopayStatus(dinopayStatus string) events.DinopayPaymentCreatedBuilder {
	b.dinopayStatus = dinopayStatus
	return b
}

func (b *DinopayPaymentCreatedBuilder) DinopayTime(dinopayTime int64) events.DinopayPaymentCreatedBuilder {
	b.dinopayTime = dinopayTime
	return b
}

func (b *DinopayPaymentCreatedBuilder) CorrelationID(correlationID string) events.DinopayPaymentCreatedBuilder {
	b.correlationID = correlationID
	return b
}

func (b *DinopayPaymentCreatedBuilder) Build() (events.DinopayPaymentCreated, error) {
	event := events.DinopayPaymentCreated{
		PaymentapiWithdrawalId: b.paymentapiWithdrawalId,
		DinopayId:              b.dinopayId,
		DinopayStatus:          b.dinopayStatus,
		DinopayTime:            b.dinopayTime,
	}
	eventJson, err := json.Marshal(event)
	if err != nil {
		return events.DinopayPaymentCreated{}, fmt.Errorf("failed building DinopayPaymentCreated event: %w", err)
	}
	cloudEvent, _ := b.cloudEventBuilder.
		Id(uuid.New().String()).
		Source(events.GatewayCloudEventsSource).
		EventType(events.EventTypeDinopayPaymentCreated).
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
