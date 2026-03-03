package realmstore

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	// "project/config"

	// "project/config"
	"reflect"
	"strings"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlserver"
	uid "github.com/dracory/uid"
)

// Store defines a session store
type Store struct {
	personTableName    string
	emailTableName     string
	db                 *sql.DB
	dbDriverName       string
	automigrateEnabled bool
	debug              bool
}

// StoreOption options for the cache store
type StoreOption func(*Store)

// WithAutoMigrate sets the table name for the cache store
func WithAutoMigrate(automigrateEnabled bool) StoreOption {
	return func(s *Store) {
		s.automigrateEnabled = automigrateEnabled
	}
}

// WithDb sets the database for the setting store
func WithDb(db *sql.DB) StoreOption {
	return func(s *Store) {
		s.db = db
		s.dbDriverName = s.DriverName(s.db)
	}
}

// WithDebug prints the SQL queries
func WithDebug(debug bool) StoreOption {
	return func(s *Store) {
		s.debug = debug
	}
}

// WithPersonTableName sets the table name for the cache store
func WithPersonTableName(personTableName string) StoreOption {
	return func(s *Store) {
		s.personTableName = personTableName
	}
}

// WithEmailTableName sets the table name for the cache store
func WithEmailTableName(emailTableName string) StoreOption {
	return func(s *Store) {
		s.emailTableName = emailTableName
	}
}

// NewStore creates a new entity store
func NewStore(opts ...StoreOption) (*Store, error) {
	store := &Store{}
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

// AutoMigrate auto migrate
func (st *Store) AutoMigrate() error {
	sqlPerson := st.SqlCreatePersonTable()

	_, errorPerson := st.db.Exec(sqlPerson)
	if errorPerson != nil {
		log.Println(errorPerson)
		return errorPerson
	}

	sqlEmail := st.SqlCreateEmailTable()

	_, errEmail := st.db.Exec(sqlEmail)
	if errEmail != nil {
		log.Println(errEmail.Error())
		return errEmail
	}

	return nil
}

// DriverName finds the driver name from database
func (st *Store) DriverName(db *sql.DB) string {
	dv := reflect.ValueOf(db.Driver())
	driverFullName := dv.Type().String()
	if strings.Contains(driverFullName, "mysql") {
		return "mysql"
	}
	if strings.Contains(driverFullName, "postgres") || strings.Contains(driverFullName, "pq") {
		return "postgres"
	}
	if strings.Contains(driverFullName, "sqlite") {
		return "sqlite"
	}
	if strings.Contains(driverFullName, "mssql") {
		return "mssql"
	}
	return driverFullName
}

// DbExec execute sql
func (st *Store) DbExec(sqlStr string) error {

	if st.debug {
		log.Println(sqlStr)
	}

	_, err := st.db.Exec(sqlStr)

	if err != nil {
		if st.debug {
			log.Println(err)
		}
		return err
	}

	return nil
}

// EnableDebug - enables the debug option
func (st *Store) EnableDebug(debug bool) {
	st.debug = debug
}

func (store *Store) PersonFindOrCreate(firstName, lastName, email string) (*Person, error) {
	person, err := store.PersonFindByEmail(email)

	if err != nil {
		return nil, err
	}

	if person != nil {
		return person, nil
	}

	person = &Person{
		Status:    PersonStatusActive,
		FirstName: firstName,
		LastName:  lastName,
	}

	err = store.PersonCreate(person)

	if err != nil {
		return nil, err
	}

	isEmailCreated, err := store.EmailCreate(&Email{
		ID:           uid.MicroUid(),
		IsDefault:    Yes,
		Status:       EmailStatusActive,
		OwnerID:      person.ID,
		EmailAddress: email,
	})

	if err != nil {
		return nil, err
	}

	if !isEmailCreated {
		return nil, errors.New("email failed to be created")
	}

	return person, nil
}
