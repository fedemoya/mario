package mario

type Event[Visitor any] interface {
	CloudEvent
	Acknowledger

	Accept(Visitor) error
}
