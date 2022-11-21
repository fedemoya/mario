package memdb

import (
	"mario"
)

type StorableCloudEvent struct {
	mario.CloudEvent

	StatusField string
}
