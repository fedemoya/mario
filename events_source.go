package mario

type RawEvent []byte

type EventsSource[EventVisitor any] interface {
	Subscribe() (<-chan RawEvent, <-chan error)
}
