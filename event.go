package mario

type Event[Visitor any] interface {
	CloudEvent

	Accept(Visitor) error
	Ack() error
	Nack(retry bool) error
}
