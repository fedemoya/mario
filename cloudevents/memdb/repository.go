package memdb

import (
	"github.com/hashicorp/go-memdb"
	"mario"
)

type CloudEvent struct {
	ID          string
	Source      string
	SpecVersion string
	Type        string
	Time        int64
	Data        []byte

	Aknowledged  bool
	DeadLettered bool
}

type Repository struct {
	db *memdb.MemDB
}

func NewRepository() *Repository {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"events": &memdb.TableSchema{
				Name: "events",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"age": &memdb.IndexSchema{
						Name:    "age",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Source"},
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

func (r Repository) Add(event mario.CloudEvent) error {
	//TODO implement me
	panic("implement me")
}
