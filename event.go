package mario

type SerializableCloudEvent interface {
	CloudEvent
	Serializer
}

type Event[Visitor any] interface {
	SerializableCloudEvent
	Acknowledger

	Accept(Visitor) error
}
