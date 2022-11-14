package mario

import "encoding/json"

type BaseEvent struct {
	cloudEventImplementor CloudEvent
	acknowledger          Acknowledger
	jsonMarshaler         json.Marshaler
}

func NewBaseEvent(cloudEventImplementor CloudEvent, acknowledger Acknowledger, jsonMarshaler json.Marshaler) BaseEvent {
	return BaseEvent{cloudEventImplementor: cloudEventImplementor, acknowledger: acknowledger, jsonMarshaler: jsonMarshaler}
}

func (be BaseEvent) ID() string {
	return be.cloudEventImplementor.ID()
}

func (be BaseEvent) Source() string {
	return be.cloudEventImplementor.Source()
}

func (be BaseEvent) Type() string {
	return be.cloudEventImplementor.Type()
}

func (be BaseEvent) Time() int64 {
	return be.cloudEventImplementor.Time()
}

func (be BaseEvent) CorrelationID() string {
	return be.cloudEventImplementor.CorrelationID()
}

func (be BaseEvent) Ack() error {
	return be.acknowledger.Ack()
}

func (be BaseEvent) Nack(retry bool) error {
	return be.acknowledger.Nack(retry)
}

func (be BaseEvent) MarshalJSON() ([]byte, error) {
	return be.jsonMarshaler.MarshalJSON()
}
