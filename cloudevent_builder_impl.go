package mario

type CloudEventBuilderImpl struct {
	cloudEventImpl *cloudEventImpl
}

func NewCloudEventBuilderImpl() *CloudEventBuilderImpl {
	cloudEventImpl := &cloudEventImpl{}
	b := &CloudEventBuilderImpl{cloudEventImpl: cloudEventImpl}
	return b
}

func (b *CloudEventBuilderImpl) Id(id string) CloudEventBuilder {
	b.cloudEventImpl.id = id
	return b
}

func (b *CloudEventBuilderImpl) Source(source string) CloudEventBuilder {
	b.cloudEventImpl.source = source
	return b
}

func (b *CloudEventBuilderImpl) SpecVersion(specVersion string) CloudEventBuilder {
	b.cloudEventImpl.specVersion = specVersion
	return b
}

func (b *CloudEventBuilderImpl) EventType(eventType string) CloudEventBuilder {
	b.cloudEventImpl.eventType = eventType
	return b
}

func (b *CloudEventBuilderImpl) Time(time int64) CloudEventBuilder {
	b.cloudEventImpl.time = time
	return b
}

func (b *CloudEventBuilderImpl) CorrelationID(correlationID string) CloudEventBuilder {
	b.cloudEventImpl.correlationID = correlationID
	return b
}

func (b *CloudEventBuilderImpl) Data(data []byte) CloudEventBuilder {
	b.cloudEventImpl.data = data
	return b
}

func (b *CloudEventBuilderImpl) Build() (CloudEvent, error) {
	return *b.cloudEventImpl, nil
}
