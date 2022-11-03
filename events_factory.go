package mario

type EventsFactory[Visitor any] interface {
	CreateEvent(event RawEvent) (AcknowledgeableEvent[Visitor], error)
}
