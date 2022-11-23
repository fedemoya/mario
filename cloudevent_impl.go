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

func (s cloudEventImpl) ID() string {
	return s.id
}

func (s cloudEventImpl) Source() string {
	return s.source
}

func (s cloudEventImpl) Type() string {
	return s.eventType
}

func (s cloudEventImpl) Time() int64 {
	return s.time
}

func (s cloudEventImpl) CorrelationID() string {
	//TODO implement me
	panic("implement me")
}

func (s cloudEventImpl) Data() []byte {
	return s.data
}

type CloudEventImplBuilder struct {
	cloudEventImpl *cloudEventImpl
}
