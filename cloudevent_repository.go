package mario

type CloudEventStatus string

const (
	CloudEventProcessed CloudEventStatus = "processed"
	CloudEventPending                    = "pending"
	CloudEventFailed                     = "failed"
)

type CloudEventRepository interface {
	Add(event CloudEvent) error
	Stream(source string) (<-chan CloudEvent, error)
	UpdateStatus(event CloudEvent, status CloudEventStatus) error
	GetProcessingRetries(cloudEvent CloudEvent) (int, error)
	IncrementRetries(cloudEvent CloudEvent) error
}
