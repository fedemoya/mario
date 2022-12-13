package json

import "encoding/json"

type CloudEvent struct {
	ID_     string          `json:"id"`
	Source_ string          `json:"source"`
	Type_   string          `json:"type"`
	Time_   int64           `json:"time"`
	Data_   json.RawMessage `json:"data"`
}

func (c CloudEvent) ID() string {
	return c.ID_
}

func (c CloudEvent) Source() string {
	return c.Source_
}

func (c CloudEvent) Type() string {
	return c.Type_
}

func (c CloudEvent) Time() int64 {
	return c.Time_
}

func (c CloudEvent) CorrelationID() string {
	return ""
}

func (c CloudEvent) Data() []byte {
	return c.Data_
}
