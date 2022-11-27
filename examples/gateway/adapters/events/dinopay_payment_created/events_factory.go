package dinopay_payment_created

import (
	"encoding/json"
	"fmt"
	"mario"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
)

type EventsFactory struct {
	acknowledger mario.Acknowledger
}

var _ mario.EventsFactory[gatewayDomainEvents.Visitor] = (*EventsFactory)(nil)

func NewEventsFactory(acknowledger mario.Acknowledger) *EventsFactory {
	return &EventsFactory{acknowledger: acknowledger}
}

func (e *EventsFactory) CreateEvent(cloudEvent mario.CloudEvent) (mario.Event[gatewayDomainEvents.Visitor], error) {
	var dinopayPaymentCreatedJSON dinopayEventCreatedJSON
	err := json.Unmarshal(cloudEvent.Data(), &dinopayPaymentCreatedJSON)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshalling dinopayPaymentCreated event with data %s: %w", cloudEvent.Data(), err)
	}
	event := gatewayDomainEvents.DinopayPaymentCreated{
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
