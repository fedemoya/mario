package mario

type Event[Visitor any] interface {
	CloudEvent

	Ack() error
	Nack(retry bool) error
	Accept(Visitor) error
}
