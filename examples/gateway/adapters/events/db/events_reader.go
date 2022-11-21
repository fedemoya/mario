package db

import (
	"context"
	"mario"
)

type EventsReader struct {
	repository   *mario.CloudEventRepository
	eventsSource string
	ch           chan mario.CloudEvent
	errCh        chan error
}

func NewCloudEventsReader(repository *mario.CloudEventRepository, eventsSource string) *EventsReader {

	return &EventsReader{
		repository:   repository,
		eventsSource: eventsSource,
	}
}

func (e *EventsReader) Subscribe() (<-chan mario.CloudEvent, <-chan error) {
	ch := make(chan mario.CloudEvent)
	errCh := make(chan error)
	e.ch = ch
	e.errCh = errCh
	return ch, errCh
}

func (e *EventsReader) Start(ctx context.Context) {
	go func() {
		for {
			select {}
		}
	}()
}
