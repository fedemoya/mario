package mario

type Acknowledger interface {
	Ack(cloudEvent CloudEvent) error
	Nack(cloudEvent CloudEvent, retry bool) error
}
