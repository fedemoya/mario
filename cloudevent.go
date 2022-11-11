package mario

import "encoding/json"

type CloudEvent interface {
	json.Marshaler

	ID() string
	Source() string
	Type() string
	Time() int64
	CorrelationID() string
}
