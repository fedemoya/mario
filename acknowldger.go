package mario

type Acknowledger interface {
	Ack() error
	Nack(retry bool) error
}
