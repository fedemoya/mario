package memdb

type Acknowledger struct {
	repository Repository
}

func (a Acknowledger) Ack() error {
	//TODO implement me
	panic("implement me")
}

func (a Acknowledger) Nack(retry bool) error {
	//TODO implement me
	panic("implement me")
}
