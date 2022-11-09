package domain

type BaseEvent struct {
	id            string
	source        string
	eventType     string
	time          int64
	correlationId string
}

func (b BaseEvent) ID() string {
	return b.id
}

func (b BaseEvent) Source() string {
	return b.source
}

func (b BaseEvent) Type() string {
	return b.eventType
}

func (b BaseEvent) Time() int64 {
	return b.time
}

func (b BaseEvent) CorrelationID() string {
	return b.correlationId
}

func (b BaseEvent) Ack() error {
	return nil
}

func (b BaseEvent) Nack(_ bool) error {
	return nil
}

func (b BaseEvent) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}
