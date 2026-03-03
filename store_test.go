package realmstore

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func newTestStore(t *testing.T) StoreInterface {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	store, err := NewStore(
		WithDb(db),
		WithPersonTableName("persons"),
		WithEmailTableName("emails"),
		WithAutoMigrate(true),
	)
	if err != nil {
		t.Fatalf("failed to create store: %v", err)
	}

	return store
}
