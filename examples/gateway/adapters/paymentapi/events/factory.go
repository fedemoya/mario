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

func (f *Factory) CreateEvent(cloudEvent mario.CloudEvent) (mario.Event[paymentapiEvents.Visitor], error) {
	switch cloudEvent.Type() {
	case "withdrawal.created":
		var withdrawalCreated WithdrawalCreated
		err := json.Unmarshal(cloudEvent.Data(), &withdrawalCreated)
		if err != nil {
			return nil, fmt.Errorf("failed unmarshalling raw evet withdrawal.created %s: %w", cloudEvent, err)
		}
	}
	return nil, nil
}
