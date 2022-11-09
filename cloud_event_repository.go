package mario

type CloudEventRepository interface {
	PersistEvent(event CloudEvent) error
}
