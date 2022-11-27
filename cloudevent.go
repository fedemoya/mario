package mario

type CloudEventStatus string

const (
	CloudEventProcessed CloudEventStatus = "processed"
	CloudEventPending                    = "pending"
	CloudEventFailed                     = "failed"
)

type CloudEvent interface {
	ID() string
	Source() string
	Type() string
	Time() int64
	CorrelationID() string
	Status() CloudEventStatus
	Data() []byte
}
