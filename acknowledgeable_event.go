package mario

type AcknowledgeableEvent[V any] interface {
	Event[V]
	Acknowledger
}
