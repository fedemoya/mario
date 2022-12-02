package events

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"mario"
	"mario/examples/gateway/domain/gateway/events"
	"time"
)

type dinopayPaymentUpdated struct {
	PaymentapiWithdrawalId string `json:"paymentapi_withdrawal_id"`
	DinopayId              string `json:"dinopay_id"`
	DinopayStatus          string `json:"dinopay_status"`
	DinopayTime            int64  `json:"dinopay_time"`
}

type DinopayPaymentUpdatedBuilder struct {
	cloudEventBuilder mario.CloudEventBuilder
	acknowledger      mario.Acknowledger

	paymentapiWithdrawalId string
	dinopayId              string
	dinopayStatus          string
	dinopayTime            int64
	correlationID          string
}

func NewDinopayPaymentUpdatedBuilder(
	cloudEventBuilder mario.CloudEventBuilder,
	acknowledger mario.Acknowledger,
) *DinopayPaymentUpdatedBuilder {

	return &DinopayPaymentUpdatedBuilder{
		cloudEventBuilder: cloudEventBuilder,
		acknowledger:      acknowledger,
	}
}

func (b *DinopayPaymentUpdatedBuilder) PaymentapiWithdrawalId(paymentapiWithdrawalId string) events.DinopayPaymentUpdatedBuilder {
	b.paymentapiWithdrawalId = paymentapiWithdrawalId
	return b
}

func (b *DinopayPaymentUpdatedBuilder) DinopayId(dinopayId string) events.DinopayPaymentUpdatedBuilder {
	b.dinopayId = dinopayId
	return b
}

func (b *DinopayPaymentUpdatedBuilder) DinopayStatus(dinopayStatus string) events.DinopayPaymentUpdatedBuilder {
	b.dinopayStatus = dinopayStatus
	return b
}

func (b *DinopayPaymentUpdatedBuilder) DinopayTime(dinopayTime int64) events.DinopayPaymentUpdatedBuilder {
	b.dinopayTime = dinopayTime
	return b
}

func (b *DinopayPaymentUpdatedBuilder) CorrelationID(correlationID string) events.DinopayPaymentUpdatedBuilder {
	b.correlationID = correlationID
	return b
}

func (b *DinopayPaymentUpdatedBuilder) Build() (events.DinopayPaymentUpdated, error) {
	event := events.DinopayPaymentUpdated{
		PaymentapiWithdrawalId: b.paymentapiWithdrawalId,
		DinopayId:              b.dinopayId,
		DinopayStatus:          b.dinopayStatus,
		DinopayTime:            b.dinopayTime,
	}
	eventJson, err := json.Marshal(event)
	if err != nil {
		return events.DinopayPaymentUpdated{}, fmt.Errorf("failed building DinopayPaymentUpdated event: %w", err)
	}
	cloudEvent, _ := b.cloudEventBuilder.
		Id(uuid.New().String()).
		Source(events.GatewayCloudEventsSource).
		EventType(events.EventTypeDinopayPaymentUpdated).
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
