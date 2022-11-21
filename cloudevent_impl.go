package mario

import "encoding/json"

type CloudEventImpl struct {
	IDField            string          `json:"id"`
	SourceField        string          `json:"source"`
	SpecVersionField   string          `json:"specversion"`
	TypeField          string          `json:"type"`
	TimeField          int64           `json:"time"`
	CorrelationIDField string          `json:"correlation_id"`
	DataField          json.RawMessage `json:"data"`
}

func (s CloudEventImpl) ID() string {
	return s.IDField
}

func (s CloudEventImpl) Source() string {
	return s.SourceField
}

func (s CloudEventImpl) Type() string {
	return s.TypeField
}

func (s CloudEventImpl) Time() int64 {
	return s.TimeField
}

func (s CloudEventImpl) CorrelationID() string {
	//TODO implement me
	panic("implement me")
}

func (s CloudEventImpl) Data() []byte {
	return s.DataField
}
