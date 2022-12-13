package events

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
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
		return f.createWithdrawalCreatedDomainEvent(cloudEvent)
	default:
		log.Debug().Msgf("ignoring event of type %s", cloudEvent.Type())
		return nil, nil
	}
}

func (f *Factory) createWithdrawalCreatedDomainEvent(cloudEvent mario.CloudEvent) (mario.Event[paymentapiEvents.Visitor], error) {
	var withdrawalCreated WithdrawalCreated
	err := json.Unmarshal(cloudEvent.Data(), &withdrawalCreated)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshalling raw evet withdrawal.created %s: %w", cloudEvent, err)
	}
	newCloudEvent, err := mario.NewCloudEventBuilderImpl().
		Id(cloudEvent.ID()).
		Source(cloudEvent.Source()).
		EventType(cloudEvent.Type()).
		Time(cloudEvent.Time()).
		CorrelationID(withdrawalCreated.Id).
		Data(cloudEvent.Data()).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed building cloud event: %w", err)
	}
	withdrawaCreatedDomainEvent := paymentapiEvents.WithdrawalCreated{
		BaseEvent: mario.NewBaseEvent(
			newCloudEvent,
			mario.DummyAcknowledger{}, // TODO implement amqp acknowledger
		),
		Id:                 withdrawalCreated.Id,
		Amount:             withdrawalCreated.Amount,
		SourceAccount:      withdrawalCreated.SourceAccount,
		DestinationAccount: withdrawalCreated.DestinationAccount,
	}
	return withdrawaCreatedDomainEvent, nil
}
