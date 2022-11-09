package mario

import "fmt"

type ErrorCallback func(err error)

type Processor[EventVisitor any] struct {
	eventsSource  EventsSource[EventVisitor]
	eventsFactory EventsFactory[EventVisitor]
	visitor       EventVisitor
	errorCb       ErrorCallback

	stop chan bool
}

func NewProcessor[EventVisitor any](
	eventsSource EventsSource[EventVisitor],
	eventsFactory EventsFactory[EventVisitor],
	visitor EventVisitor,
	errorCb ErrorCallback,
) *Processor[EventVisitor] {

	return &Processor[EventVisitor]{
		eventsSource:  eventsSource,
		eventsFactory: eventsFactory,
		visitor:       visitor,
		errorCb:       errorCb,

		stop: make(chan bool),
	}
}

func (p *Processor[V]) Start() error {
	eventsCh, errorsCh := p.eventsSource.Subscribe()
	go func() {
		for {
			select {
			case rawEvent := <-eventsCh:
				go p.processEvent(rawEvent)
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

func (p *Processor[V]) processEvent(rawEvent RawEvent) error {
	event, err := p.eventsFactory.CreateEvent(rawEvent)
	if err != nil {
		return fmt.Errorf("failed creating event from rawEvent %s", rawEvent)
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
