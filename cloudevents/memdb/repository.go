package memdb

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"mario"
)

const (
	StorableEventStatusNotProcessed     = "not_processed"
	StorableEventStatusProcessed        = "processed"
	StorableEventStatusProcessingFailed = "processing_failed"
)

type StorableEvent struct {
	ID          string
	Source      string
	SpecVersion string
	Type        string
	Time        int64
	Data        []byte

	Status string
}

type Repository struct {
	db *memdb.MemDB
}

func NewRepository() *Repository {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"events": {
				Name: "events",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"source": {
						Name:    "source",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Source"},
					},
					"status": {
						Name:    "status",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Status"},
					},
				},
			},
		},
	}
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	return &Repository{db}
}

func (r Repository) Add(event mario.SerializableCloudEvent) error {
	data, err := event.Serialize()
	if err != nil {
		return fmt.Errorf("failed adding event with source %s and type %s to repository: %w",
			event.Source(),
			event.Type(),
			err,
		)
	}
	storableCloudEvent := StorableEvent{
		ID:          event.ID(),
		Source:      event.Source(),
		SpecVersion: "",
		Type:        event.Type(),
		Time:        event.Time(),
		Data:        data,
		Status:      StorableEventStatusNotProcessed,
	}
	txn := r.db.Txn(true)
	err = txn.Insert("events", storableCloudEvent)
	if err != nil {
		return fmt.Errorf("failed adding event with source %s and type %s to repository: %w",
			event.Source(),
			event.Type(),
			err)
	}
	return nil
}
