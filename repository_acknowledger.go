package mario

import (
	"fmt"
)

type RepositoryAcknowledger struct {
	repository CloudEventRepository
	maxRetries int
}

func NewRepositoryAcknowledger(repository CloudEventRepository, maxRetries int) *RepositoryAcknowledger {
	return &RepositoryAcknowledger{repository: repository, maxRetries: maxRetries}
}

func (a RepositoryAcknowledger) Ack(cloudEvent CloudEvent) error {
	// TODO check if cloudEvent is not in failed state ?
	err := a.repository.UpdateStatus(cloudEvent, CloudEventProcessed)
	if err != nil {
		return fmt.Errorf("failed acknowledging cloudEvent with ID %s: %w", cloudEvent.ID(), err)
	}
	return nil
}

func (a RepositoryAcknowledger) Nack(cloudEvent CloudEvent, retry bool) error {
	if retry {
		// TODO check if cloudEvent is in failed state
		retries, err := a.repository.GetProcessingRetries(cloudEvent)
		if err != nil {
			return fmt.Errorf("failed nacknowledging cloudEvent with ID %s: %w", cloudEvent.ID(), err)
		}
		if retries <= a.maxRetries {
			err := a.repository.IncrementRetries(cloudEvent)
			if err != nil {
				return fmt.Errorf("failed nacknowledging cloudEvent with ID %s: %w", cloudEvent.ID(), err)
			}
			return nil
		}
	}
	err := a.repository.UpdateStatus(cloudEvent, CloudEventFailed)
	if err != nil {
		return fmt.Errorf("failed nacknowledging cloudEvent with ID %s: %w", cloudEvent.ID(), err)
	}
	return nil
}
