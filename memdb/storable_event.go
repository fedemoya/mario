package memdb

type StorableCloudEvent struct {
	ID            string
	Source        string
	Type          string
	Time          int64
	CorrelationID string
	Data          []byte
	Status        string
}
