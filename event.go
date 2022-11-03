package mario

type Event[Visitor any] interface {
	Accept(Visitor) error
	Ack() error
	Nack(opts interface{}) error
}
