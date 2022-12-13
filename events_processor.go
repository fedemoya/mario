package mario

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ErrorCallback func(err error)

type Processor[EventVisitor any] struct {
	eventsReader  EventsReader
	eventsFactory EventsFactory[EventVisitor]
	visitor       EventVisitor
	errorCb       ErrorCallback

	stop chan bool
}

func NewProcessor[EventVisitor any](
	eventsReader EventsReader,
	eventsFactory EventsFactory[EventVisitor],
	visitor EventVisitor,
	errorCb ErrorCallback,
) *Processor[EventVisitor] {

	return &Processor[EventVisitor]{
		eventsReader:  eventsReader,
		eventsFactory: eventsFactory,
		visitor:       visitor,
		errorCb:       errorCb,

		stop: make(chan bool),
	}
}

func (p *Processor[V]) Start() error {
	eventsCh, errorsCh := p.eventsReader.Subscribe()
	go func() {
		for {
			select {
			case cloudEvent := <-eventsCh:
				go p.processEvent(cloudEvent)
			case handlerErr := <-errorsCh:
				go p.errorCb(handlerErr)
			case <-p.stop:
				return
			}
		}
	}()
	return nil
}

func (p *Processor[V]) Stop() error {
	close(p.stop)
	return nil
}

func (p *Processor[V]) processEvent(cloudEvent CloudEvent) {
	event, err := p.eventsFactory.CreateEvent(cloudEvent)
	if err != nil {
		p.errorCb(fmt.Errorf("failed creating event from cloudEvent %s: %w", cloudEvent, err))
		return
	}
	err = event.Accept(p.visitor)
	if err != nil {
		logger := p.eventErrorLogger(event, err)
		_, retryable := err.(IsRetryableError)
		if retryable {
			logger.Error().Msgf("failed processing event with retryable error")
			event.Nack(true)
		} else {
			logger.Error().Msgf("failed processing event with non-retryable error")
			event.Nack(false)
		}
	} else {
		event.Ack()
	}
}

func (p *Processor[V]) eventErrorLogger(cloudEvent CloudEvent, err error) zerolog.Logger {
	return log.With().
		Str("eventId", cloudEvent.ID()).
		Str("eventType", cloudEvent.Type()).
		Str("eventSource", cloudEvent.Source()).
		Err(err).
		Logger()
}
