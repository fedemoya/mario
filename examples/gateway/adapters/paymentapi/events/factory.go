package events

import (
	"encoding/json"
	"fmt"
	"mario"
	paymentapiEvents "mario/examples/gateway/domain/paymentapi/events"
)

type Factory struct {
}

var _ mario.EventsFactory[paymentapiEvents.Visitor] = (*Factory)(nil)

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) CreateEvent(event mario.RawEvent) (mario.Event[paymentapiEvents.Visitor], error) {
	var cloudevent mario.CloudEvent
	err := json.Unmarshal(event, &cloudevent)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshalling raw cloudevent %s: %w", event, err)
	}
	switch cloudevent.Type {
	case "withdrawal.created":
		var withdrawalCreated WithdrawalCreated
		err := json.Unmarshal(cloudevent.Data, &withdrawalCreated)
		if err != nil {
			return nil, fmt.Errorf("failed unmarshalling raw evet withdrawal.created %s: %w", event, err)
		}
	}
	return nil, nil
}
