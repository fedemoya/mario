package mario

type CloudEventRepository interface {
	Add(event CloudEvent) error
	Stream(source string) (<-chan CloudEvent, error)
	UpdateStatus(event CloudEvent, status CloudEventStatus) error
}
