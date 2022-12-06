package mario

type BaseEvent struct {
	cloudEvent   CloudEvent
	acknowledger Acknowledger
}

func NewBaseEvent(cloudEventImplementor CloudEvent, acknowledger Acknowledger) BaseEvent {
	return BaseEvent{cloudEvent: cloudEventImplementor, acknowledger: acknowledger}
}

func (be BaseEvent) ID() string {
	return be.cloudEvent.ID()
}

func (be BaseEvent) Source() string {
	return be.cloudEvent.Source()
}

func (be BaseEvent) Type() string {
	return be.cloudEvent.Type()
}

func (be BaseEvent) Time() int64 {
	return be.cloudEvent.Time()
}

func (be BaseEvent) CorrelationID() string {
	return be.cloudEvent.CorrelationID()
}

func (be BaseEvent) Data() []byte {
	return be.cloudEvent.Data()
}

func (be BaseEvent) Ack() error {
	return be.acknowledger.Ack(be.cloudEvent)
}

func (be BaseEvent) Nack(retry bool) error {
	return be.acknowledger.Nack(be.cloudEvent, retry)
}
