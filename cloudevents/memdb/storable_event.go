package memdb

import "mario"

type StorableCloudEvent struct {
	ID                string
	Source            string
	Type              string
	Time              int64
	CorrelationID     string
	Data              []byte
	Status            mario.CloudEventStatus
	ProcessingRetries int
}
