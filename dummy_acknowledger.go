package mario

type DummyAcknowledger struct {
}

func (d DummyAcknowledger) Ack(_ CloudEvent) error {
	return nil
}

func (d DummyAcknowledger) Nack(_ CloudEvent, _ bool) error {
	return nil
}
