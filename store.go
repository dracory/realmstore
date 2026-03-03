package realmstore

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrEmailCreationFailed = errors.New("email failed to be created")
)

type StoreOption func(*storeImplementation)

func WithAutoMigrate(automigrateEnabled bool) StoreOption {
	return func(s *storeImplementation) {
		s.automigrateEnabled = automigrateEnabled
	}
}

func WithDb(db *sql.DB) StoreOption {
	return func(s *storeImplementation) {
		s.db = db
		s.dbDriverName = s.DriverName(s.db)
	}
}

func WithDebug(debug bool) StoreOption {
	return func(s *storeImplementation) {
		s.debug = debug
	}
}

func WithPersonTableName(personTableName string) StoreOption {
	return func(s *storeImplementation) {
		s.personTableName = personTableName
	}
}

func WithEmailTableName(emailTableName string) StoreOption {
	return func(s *storeImplementation) {
		s.emailTableName = emailTableName
	}
}

func NewStore(opts ...StoreOption) (StoreInterface, error) {
	store := &storeImplementation{}
	for _, opt := range opts {
		opt(store)
	}

	if store.personTableName == "" {
		return nil, errors.New("realm store: personTableName is required")
	}

	if store.emailTableName == "" {
		return nil, errors.New("realm store: emailTableName is required")
	}

	if store.automigrateEnabled {
		if err := store.AutoMigrate(); err != nil {
			return nil, fmt.Errorf("failed to auto-migrate: %w", err)
		}
	}

	return store, nil
}
