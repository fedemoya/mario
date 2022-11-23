package mario

type CloudEventRepository interface {
	Add(event CloudEvent) error
	Stream(source string) (<-chan CloudEvent, error)
	Ack(event CloudEvent) error
	Nack(event CloudEvent, retry bool) error
}
