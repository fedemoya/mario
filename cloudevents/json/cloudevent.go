package json

import "encoding/json"

type CloudEvent struct {
	ID          string          `json:"id"`
	Source      string          `json:"source"`
	SpecVersion string          `json:"specversion"`
	Type        string          `json:"type"`
	Time        int64           `json:"time"`
	Data        json.RawMessage `json:"data"`
}
