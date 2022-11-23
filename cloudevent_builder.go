package mario

type CloudEventBuilder interface {
	Id(id string) CloudEventBuilder
	Source(source string) CloudEventBuilder
	SpecVersion(specVersion string) CloudEventBuilder
	EventType(eventType string) CloudEventBuilder
	Time(time int64) CloudEventBuilder
	CorrelationID(correlationID string) CloudEventBuilder
	Data(data []byte) CloudEventBuilder
	Build() (CloudEvent, error)
}
