package mario

type cloudEventImpl struct {
	id            string
	source        string
	specVersion   string
	eventType     string
	time          int64
	correlationID string
	data          []byte
}

var _ CloudEvent = cloudEventImpl{}

func (ce cloudEventImpl) ID() string {
	return ce.id
}

func (ce cloudEventImpl) Source() string {
	return ce.source
}

func (ce cloudEventImpl) Type() string {
	return ce.eventType
}

func (ce cloudEventImpl) Time() int64 {
	return ce.time
}

func (ce cloudEventImpl) CorrelationID() string {
	return ce.correlationID
}

func (ce cloudEventImpl) Data() []byte {
	return ce.data
}
