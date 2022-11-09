package events

import (
	"mario"
	paymentapiEvents "mario/examples/gateway/domain/paymentapi/events"
)

type Factory struct {
}

func (f Factory) CreateEvent(event mario.RawEvent) (mario.Event[paymentapiEvents.Visitor], error) {
	//TODO implement me
	panic("implement me")
}
