package mario

type EventsFactory[Visitor any] interface {
	CreateEvent(event RawEvent) (Event[Visitor], error)
}
