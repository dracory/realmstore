package realmstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/dracory/neat"
	contractsorm "github.com/dracory/neat/contracts/database/orm"
	contractsschema "github.com/dracory/neat/contracts/database/schema"
	"github.com/dromara/carbon/v2"
)

// == INTERFACE ===============================================================

type StoreInterface interface {
	// GetDB returns the underlying *sql.DB
	GetDB() *sql.DB
	// GetPersonTableName returns the person table name
	GetPersonTableName() string
	// SetPersonTableName sets the person table name
	SetPersonTableName(personTableName string)
	// GetEmailTableName returns the email table name
	GetEmailTableName() string
	// SetEmailTableName sets the email table name
	SetEmailTableName(emailTableName string)

	// MigrateDown drops the realm store tables
	MigrateDown(ctx context.Context, tx ...*sql.Tx) error
	// MigrateUp creates the realm store tables
	MigrateUp(ctx context.Context, tx ...*sql.Tx) error

	EnableDebug(debug bool)

	PersonCreate(ctx context.Context, person PersonInterface) error
	PersonCreateWithEmail(ctx context.Context, fields map[string]string) (PersonInterface, error)
	PersonFindOrCreate(ctx context.Context, firstName, lastName, email string) (PersonInterface, error)
	PersonFindByID(ctx context.Context, personID string) (PersonInterface, error)
	PersonFindByEmail(ctx context.Context, email string) (PersonInterface, error)
	PersonFindDefaultEmail(ctx context.Context, personID string) (EmailInterface, error)
	PersonDelete(ctx context.Context, personID string) error
	PersonList(ctx context.Context, query PersonQueryInterface) ([]PersonWithEmailInterface, error)
	PersonCount(ctx context.Context, query PersonQueryInterface) (int64, error)
	PersonMaxPages(ctx context.Context, query PersonQueryInterface) (uint, error)
	PersonUpdate(ctx context.Context, person PersonInterface) error
	PersonUpdateByID(ctx context.Context, personID string, fields map[string]any) error

	EmailCreate(ctx context.Context, email EmailInterface) error
	EmailDeleteByID(ctx context.Context, emailID string) error
	EmailFind(ctx context.Context, query EmailQueryInterface) (EmailInterface, error)
	EmailFindByID(ctx context.Context, id string) (EmailInterface, error)
	EmailFindByEmailAddress(ctx context.Context, emailAddress string) (EmailInterface, error)
	EmailFindByPersonID(ctx context.Context, personID string) (EmailInterface, error)
	EmailList(ctx context.Context, query EmailQueryInterface) ([]EmailInterface, error)
	EmailCount(ctx context.Context, query EmailQueryInterface) (int64, error)
	EmailUpdate(ctx context.Context, email EmailInterface) error
	EmailUpdateByID(ctx context.Context, emailID string, fields map[string]any) error
}

// == TYPE ====================================================================

var _ StoreInterface = (*storeImplementation)(nil)

type storeImplementation struct {
	personTableName    string
	emailTableName     string
	db                 *neat.Database
	automigrateEnabled bool
	debugEnabled       bool
	logger             *slog.Logger
}

// NewStoreOptions defines the options for creating a new realm store.
type NewStoreOptions struct {
	PersonTableName    string
	EmailTableName     string
	DB                 *sql.DB
	AutomigrateEnabled bool
	DebugEnabled       bool
}

// NewStore creates a new realm store.
func NewStore(opts NewStoreOptions) (StoreInterface, error) {
	if opts.DB == nil {
		return nil, errors.New("realm store: DB is required")
	}

	if opts.PersonTableName == "" {
		return nil, errors.New("realm store: personTableName is required")
	}

	if opts.EmailTableName == "" {
		return nil, errors.New("realm store: emailTableName is required")
	}

	neatDB, err := neat.NewFromSQLDB(opts.DB)
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	store := &storeImplementation{
		personTableName:    opts.PersonTableName,
		emailTableName:     opts.EmailTableName,
		db:                 neatDB,
		automigrateEnabled: opts.AutomigrateEnabled,
		debugEnabled:       opts.DebugEnabled,
		logger:             logger,
	}

	if store.automigrateEnabled {
		if err := store.MigrateUp(context.Background()); err != nil {
			return nil, fmt.Errorf("failed to auto-migrate: %w", err)
		}
	}

	return store, nil
}

// == DB ======================================================================

// GetDB returns the underlying *sql.DB.
func (st *storeImplementation) GetDB() *sql.DB {
	db, _ := st.db.DB()
	return db
}

// == TABLE NAMES =============================================================

// GetPersonTableName returns the person table name.
func (st *storeImplementation) GetPersonTableName() string {
	return st.personTableName
}

// SetPersonTableName sets the person table name.
func (st *storeImplementation) SetPersonTableName(personTableName string) {
	st.personTableName = personTableName
}

// GetEmailTableName returns the email table name.
func (st *storeImplementation) GetEmailTableName() string {
	return st.emailTableName
}

// SetEmailTableName sets the email table name.
func (st *storeImplementation) SetEmailTableName(emailTableName string) {
	st.emailTableName = emailTableName
}

// == MIGRATE =================================================================

// MigrateUp creates the realm store tables if they do not already exist.
func (st *storeImplementation) MigrateUp(ctx context.Context, tx ...*sql.Tx) error {
	if st.db.Schema().HasTable(st.personTableName) {
		if st.debugEnabled {
			st.logger.Info("MigrateUp: person table already exists", "table", st.personTableName)
		}
	} else {
		err := st.db.Schema().Create(st.personTableName, func(table contractsschema.Blueprint) {
			table.String(COLUMN_ID, 40)
			table.Primary(COLUMN_ID)
			table.String(COLUMN_STATUS, 40)
			table.String(COLUMN_FIRST_NAME, 100)
			table.String(COLUMN_MIDDLE_NAMES, 100).Nullable()
			table.String(COLUMN_LAST_NAME, 100)
			table.DateTime(COLUMN_BIRTH_DAY).Nullable()
			table.DateTime(COLUMN_CREATED_AT)
			table.DateTime(COLUMN_UPDATED_AT)
			table.DateTime(COLUMN_SOFT_DELETED_AT)
		})
		if err != nil {
			if st.debugEnabled {
				st.logger.Error("MigrateUp failed for person table", "error", err)
			}
			return err
		}
	}

	if st.db.Schema().HasTable(st.emailTableName) {
		if st.debugEnabled {
			st.logger.Info("MigrateUp: email table already exists", "table", st.emailTableName)
		}
	} else {
		err := st.db.Schema().Create(st.emailTableName, func(table contractsschema.Blueprint) {
			table.String(COLUMN_ID, 40)
			table.Primary(COLUMN_ID)
			table.String(COLUMN_STATUS, 40)
			table.String(COLUMN_OWNER_ID, 40)
			table.String(COLUMN_EMAIL_ADDRESS, 100)
			table.String(COLUMN_IS_DEFAULT, 3)
			table.DateTime(COLUMN_CREATED_AT)
			table.DateTime(COLUMN_UPDATED_AT)
			table.DateTime(COLUMN_SOFT_DELETED_AT)
		})
		if err != nil {
			if st.debugEnabled {
				st.logger.Error("MigrateUp failed for email table", "error", err)
			}
			return err
		}
	}

	return nil
}

// MigrateDown drops the realm store tables.
func (st *storeImplementation) MigrateDown(ctx context.Context, tx ...*sql.Tx) error {
	if st.db.Schema().HasTable(st.emailTableName) {
		err := st.db.Schema().Drop(st.emailTableName)
		if err != nil {
			if st.debugEnabled {
				st.logger.Error("MigrateDown failed for email table", "error", err)
			}
			return err
		}
	}

	if st.db.Schema().HasTable(st.personTableName) {
		err := st.db.Schema().Drop(st.personTableName)
		if err != nil {
			if st.debugEnabled {
				st.logger.Error("MigrateDown failed for person table", "error", err)
			}
			return err
		}
	}

	return nil
}

// == DEBUG ===================================================================

// EnableDebug enables or disables debug mode.
func (st *storeImplementation) EnableDebug(debug bool) {
	st.debugEnabled = debug
	if debug {
		st.db.EnableDebug()
		st.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	} else {
		st.db.DisableDebug()
		st.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
}

// == QUERY BUILDERS ===========================================================

func (st *storeImplementation) buildPersonQuery(query PersonQueryInterface) contractsorm.Query {
	q := st.db.Query().Table(st.personTableName)

	if query == nil {
		return q
	}

	if query.IsIDSet() && query.GetID() != "" {
		q = q.Where(COLUMN_ID+" = ?", query.GetID())
	}

	if query.IsStatusSet() && query.GetStatus() != "" {
		q = q.Where(COLUMN_STATUS+" = ?", query.GetStatus())
	}

	if query.IsSearchSet() && query.GetSearch() != "" {
		search := "%" + query.GetSearch() + "%"
		q = q.Where("("+COLUMN_FIRST_NAME+" LIKE ? OR "+COLUMN_LAST_NAME+" LIKE ? OR "+COLUMN_EMAIL_ADDRESS+" LIKE ?)", search, search, search)
	}

	if query.IsDeletedOnlySet() && query.GetDeletedOnly() {
		q = q.Where(COLUMN_SOFT_DELETED_AT+" < ?", carbon.Parse(MAX_DATETIME, carbon.UTC).StdTime())
	} else if !query.IsWithSoftDeletedSet() || !query.GetWithSoftDeleted() {
		q = q.Where(COLUMN_SOFT_DELETED_AT+" = ?", carbon.Parse(MAX_DATETIME, carbon.UTC).StdTime())
	}

	if query.IsLimitSet() && query.GetLimit() > 0 {
		q = q.Limit(query.GetLimit())
	}

	if query.IsOffsetSet() && query.GetOffset() > 0 {
		q = q.Offset(query.GetOffset())
	}

	if query.IsSortBySet() && query.GetSortBy() != "" {
		sortOrder := "desc"
		if query.IsSortOrderSet() && query.GetSortOrder() != "" {
			sortOrder = query.GetSortOrder()
		}
		q = q.OrderBy(query.GetSortBy(), sortOrder)
	}

	return q
}

func (st *storeImplementation) buildEmailQuery(query EmailQueryInterface) contractsorm.Query {
	q := st.db.Query().Table(st.emailTableName)

	if query == nil {
		return q
	}

	if query.IsIDSet() && query.GetID() != "" {
		q = q.Where(COLUMN_ID+" = ?", query.GetID())
	}

	if query.IsOwnerIDSet() && query.GetOwnerID() != "" {
		q = q.Where(COLUMN_OWNER_ID+" = ?", query.GetOwnerID())
	}

	if query.IsStatusSet() && query.GetStatus() != "" {
		q = q.Where(COLUMN_STATUS+" = ?", query.GetStatus())
	}

	if query.IsEmailAddressSet() && query.GetEmailAddress() != "" {
		q = q.Where(COLUMN_EMAIL_ADDRESS+" = ?", query.GetEmailAddress())
	}

	if query.IsIsDefaultSet() && query.GetIsDefault() != "" {
		q = q.Where(COLUMN_IS_DEFAULT+" = ?", query.GetIsDefault())
	}

	if query.IsSearchSet() && query.GetSearch() != "" {
		search := "%" + query.GetSearch() + "%"
		q = q.Where(COLUMN_EMAIL_ADDRESS+" LIKE ?", search)
	}

	if query.IsDeletedOnlySet() && query.GetDeletedOnly() {
		q = q.Where(COLUMN_SOFT_DELETED_AT+" < ?", carbon.Parse(MAX_DATETIME, carbon.UTC).StdTime())
	} else if !query.IsWithSoftDeletedSet() || !query.GetWithSoftDeleted() {
		q = q.Where(COLUMN_SOFT_DELETED_AT+" = ?", carbon.Parse(MAX_DATETIME, carbon.UTC).StdTime())
	}

	if query.IsLimitSet() && query.GetLimit() > 0 {
		q = q.Limit(query.GetLimit())
	}

	if query.IsOffsetSet() && query.GetOffset() > 0 {
		q = q.Offset(query.GetOffset())
	}

	if query.IsSortBySet() && query.GetSortBy() != "" {
		sortOrder := "desc"
		if query.IsSortOrderSet() && query.GetSortOrder() != "" {
			sortOrder = query.GetSortOrder()
		}
		q = q.OrderBy(query.GetSortBy(), sortOrder)
	}

	return q
}
