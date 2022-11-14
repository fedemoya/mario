package mario

type DummyAcknowledger struct {
}

func (d DummyAcknowledger) Ack() error {
	return nil
}

func (d DummyAcknowledger) Nack(_ bool) error {
	return nil
}
