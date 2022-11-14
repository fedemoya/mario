package mario

type CloudEventRepository interface {
	Add(event CloudEvent) error
}
