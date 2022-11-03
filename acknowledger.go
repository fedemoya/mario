package mario

type Acknowledger interface {
	Ack() error
	Nack(opts interface{}) error
}
