package data

import (
	"github.com/assefamaru/cellosaurus-api/pkg/db"
)

// ObjectStore provides an API for interacting
// with the data store that contains the
// ingested Cellosaurus data.
type ObjectStore struct {
	client *db.MySQLClient
}

// NewObjectStore creates a new ObjectStore.
func NewObjectStore(client *db.MySQLClient) *ObjectStore {
	return &ObjectStore{
		client: client,
	}
}
