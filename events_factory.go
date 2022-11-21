package mario

type EventsFactory[Visitor any] interface {
	CreateEvent(event CloudEvent) (Event[Visitor], error)
}
