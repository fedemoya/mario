package mario

type CloudEventRepository interface {
	Persist(event CloudEvent) error
}
