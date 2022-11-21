package mario

type EventsReader interface {
	Subscribe() (<-chan CloudEvent, <-chan error)
}
