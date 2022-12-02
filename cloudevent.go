package mario

type CloudEvent interface {
	ID() string
	Source() string
	Type() string
	Time() int64
	CorrelationID() string
	Data() []byte
}
