package events

import (
	"fmt"
	"mario"
)

type PersistenceVisitor struct {
	gatewayEventsRepository mario.CloudEventRepository
}

func (d PersistenceVisitor) VisitDinopayPaymentCreated(created DinopayPaymentCreated) error {
	return d.updateWithdrawal(created)
}

func (d PersistenceVisitor) VisitDinopayPaymentUpdated(updated DinopayPaymentUpdated) error {
	return d.updateWithdrawal(updated)
}

func (d PersistenceVisitor) updateWithdrawal(event mario.CloudEvent) error {
	err := d.gatewayEventsRepository.Persist(event)
	if err != nil {
		return fmt.Errorf("failed persisting event %s with id %s: %w", event.Type(), event.ID(), err)
	}
	return nil
}
