package memdb

import "github.com/hashicorp/go-memdb"

func InitDB() *memdb.MemDB {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"events": {
				Name: "events",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "IDField"},
					},
					"source": {
						Name:    "source",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SourceField"},
					},
					"status": {
						Name:    "status",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "StatusField"},
					},
				},
			},
		},
	}
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	return db
}
