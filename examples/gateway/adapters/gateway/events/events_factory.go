package events

import (
	"encoding/json"
	"fmt"
	"mario"
	"mario/examples/gateway/domain/gateway/events"
)

type EventsFactory struct {
	acknowledger mario.Acknowledger
}

var _ mario.EventsFactory[events.Visitor] = (*EventsFactory)(nil)

func NewEventsFactory(acknowledger mario.Acknowledger) *EventsFactory {
	return &EventsFactory{acknowledger: acknowledger}
}

func (e *EventsFactory) CreateEvent(cloudEvent mario.CloudEvent) (mario.Event[events.Visitor], error) {
	var dinopayPaymentCreatedJSON dinopayPaymentCreated
	err := json.Unmarshal(cloudEvent.Data(), &dinopayPaymentCreatedJSON)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshalling dinopayPaymentCreated event with data %s: %w", cloudEvent.Data(), err)
	}
	event := events.DinopayPaymentCreated{
		PaymentapiWithdrawalId: dinopayPaymentCreatedJSON.PaymentapiWithdrawalId,
		DinopayId:              dinopayPaymentCreatedJSON.DinopayId,
		DinopayStatus:          dinopayPaymentCreatedJSON.DinopayStatus,
		DinopayTime:            dinopayPaymentCreatedJSON.DinopayTime,
	}
	baseEvent := mario.NewBaseEvent(
		cloudEvent,
		e.acknowledger,
	)
	event.BaseEvent = baseEvent
	return event, nil
}
