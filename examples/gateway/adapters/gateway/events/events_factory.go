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
	switch cloudEvent.Type() {
	case events.EventTypeDinopayPaymentCreated:
		var dinopayPaymentCreatedWithJsonTags dinopayPaymentCreated
		err := json.Unmarshal(cloudEvent.Data(), &dinopayPaymentCreatedWithJsonTags)
		if err != nil {
			return nil, fmt.Errorf("failed unmarshalling dinopayPaymentCreated event with data %s: %w", cloudEvent.Data(), err)
		}
		dinopayPaymentCreated := events.DinopayPaymentCreated{
			PaymentapiWithdrawalId: dinopayPaymentCreatedWithJsonTags.PaymentapiWithdrawalId,
			DinopayId:              dinopayPaymentCreatedWithJsonTags.DinopayId,
			DinopayStatus:          dinopayPaymentCreatedWithJsonTags.DinopayStatus,
			DinopayTime:            dinopayPaymentCreatedWithJsonTags.DinopayTime,
		}
		baseEvent := mario.NewBaseEvent(
			cloudEvent,
			e.acknowledger,
		)
		dinopayPaymentCreated.BaseEvent = baseEvent
		return dinopayPaymentCreated, nil
	case events.EventTypeDinopayPaymentUpdated:
		var dinopayPaymentCreatedJsonTags dinopayPaymentUpdated
		err := json.Unmarshal(cloudEvent.Data(), &dinopayPaymentCreatedJsonTags)
		if err != nil {
			return nil, fmt.Errorf("failed unmarshalling dinopayPaymentCreated event with data %s: %w", cloudEvent.Data(), err)
		}
		dinopayPaymentCreated := events.DinopayPaymentUpdated{
			PaymentapiWithdrawalId: dinopayPaymentCreatedJsonTags.PaymentapiWithdrawalId,
			DinopayId:              dinopayPaymentCreatedJsonTags.DinopayId,
			DinopayStatus:          dinopayPaymentCreatedJsonTags.DinopayStatus,
			DinopayTime:            dinopayPaymentCreatedJsonTags.DinopayTime,
		}
		baseEvent := mario.NewBaseEvent(
			cloudEvent,
			e.acknowledger,
		)
		dinopayPaymentCreated.BaseEvent = baseEvent
		return dinopayPaymentCreated, nil
	default:
		return nil, nil
	}
}
