package mario

import (
	"context"
	"fmt"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
)

type RepositoryEventsReader struct {
	repository   CloudEventRepository
	eventsSource string
	ch           chan CloudEvent
	errCh        chan error
}

func NewCloudEventsReader(repository CloudEventRepository, eventsSource string) *RepositoryEventsReader {

	return &RepositoryEventsReader{
		repository:   repository,
		eventsSource: eventsSource,
	}
}

func (e *RepositoryEventsReader) Subscribe() (<-chan CloudEvent, <-chan error) {
	ch := make(chan CloudEvent)
	errCh := make(chan error)
	e.ch = ch
	e.errCh = errCh
	return ch, errCh
}

func (e *RepositoryEventsReader) Start(ctx context.Context) error {
	cloudEventsCh, err := e.repository.Stream(gatewayDomainEvents.GatewayCloudEventsSource)
	if err != nil {
		return fmt.Errorf("failed streaming cloud events from repository: %w", err)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case cloudEvent := <-cloudEventsCh:
				e.ch <- cloudEvent
			}
		}
	}()
	return nil
}
