package mario

import "fmt"

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

func (p *Processor[V]) processEvent(cloudEvent CloudEvent) error {
	event, err := p.eventsFactory.CreateEvent(cloudEvent)
	if err != nil {
		return fmt.Errorf("failed creating event from cloudEvent %s", cloudEvent)
	}
	err = event.Accept(p.visitor)
	if err != nil {
		_, retryable := err.(IsRetryableError)
		if retryable {
			event.Nack(true)
		} else {
			event.Nack(false)
		}
	} else {
		event.Ack()
	}
	return nil
}
