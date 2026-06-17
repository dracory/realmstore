package realmstore

import (
	"context"
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func newTestStore(t *testing.T) StoreInterface {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	store, err := NewStore(NewStoreOptions{
		DB:                 db,
		PersonTableName:    "persons",
		EmailTableName:     "emails",
		AutomigrateEnabled: true,
	})
	if err != nil {
		t.Fatalf("failed to create store: %v", err)
	}

	return store
}

func TestNewStoreValidation(t *testing.T) {
	t.Run("nil DB", func(t *testing.T) {
		_, err := NewStore(NewStoreOptions{
			PersonTableName: "persons",
			EmailTableName:  "emails",
		})
		if err == nil {
			t.Fatal("expected error for nil DB")
		}
	})

	t.Run("empty person table name", func(t *testing.T) {
		db, _ := sql.Open("sqlite", ":memory:")
		_, err := NewStore(NewStoreOptions{
			DB:             db,
			EmailTableName: "emails",
		})
		if err == nil {
			t.Fatal("expected error for empty person table name")
		}
	})

	t.Run("empty email table name", func(t *testing.T) {
		db, _ := sql.Open("sqlite", ":memory:")
		_, err := NewStore(NewStoreOptions{
			DB:              db,
			PersonTableName: "persons",
		})
		if err == nil {
			t.Fatal("expected error for empty email table name")
		}
	})
}

func TestMigrateUpAndDown(t *testing.T) {
	store := newTestStore(t)

	err := store.MigrateDown(context.Background())
	if err != nil {
		t.Fatalf("failed to migrate down: %v", err)
	}

	err = store.MigrateUp(context.Background())
	if err != nil {
		t.Fatalf("failed to migrate up: %v", err)
	}
}

func TestEnableDebug(t *testing.T) {
	store := newTestStore(t)

	store.EnableDebug(true)
	store.EnableDebug(false)
}

func TestTableNames(t *testing.T) {
	store := newTestStore(t)

	if store.GetPersonTableName() != "persons" {
		t.Errorf("expected person table name 'persons', got '%s'", store.GetPersonTableName())
	}

	if store.GetEmailTableName() != "emails" {
		t.Errorf("expected email table name 'emails', got '%s'", store.GetEmailTableName())
	}

	store.SetPersonTableName("people")
	if store.GetPersonTableName() != "people" {
		t.Errorf("expected person table name 'people', got '%s'", store.GetPersonTableName())
	}

	store.SetEmailTableName("mail")
	if store.GetEmailTableName() != "mail" {
		t.Errorf("expected email table name 'mail', got '%s'", store.GetEmailTableName())
	}
}

func TestGetDB(t *testing.T) {
	store := newTestStore(t)

	db := store.GetDB()
	if db == nil {
		t.Fatal("expected non-nil DB")
	}
}
