package realmstore

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlserver"
	uid "github.com/dracory/uid"
)

type storeImplementation struct {
	personTableName    string
	emailTableName     string
	db                 *sql.DB
	dbDriverName       string
	automigrateEnabled bool
	debug              bool
}

// MigrateUp creates the realm store tables
func (st *storeImplementation) MigrateUp(ctx context.Context, tx ...*sql.Tx) error {
	var txToUse *sql.Tx
	if len(tx) > 0 {
		txToUse = tx[0]
	}

	sqlPerson := st.SqlCreatePersonTable()

	var errExec error
	if txToUse != nil {
		_, errExec = txToUse.ExecContext(ctx, sqlPerson)
	} else {
		_, errExec = st.db.ExecContext(ctx, sqlPerson)
	}
	if errExec != nil {
		return errExec
	}

	sqlEmail := st.SqlCreateEmailTable()

	if txToUse != nil {
		_, errExec = txToUse.ExecContext(ctx, sqlEmail)
	} else {
		_, errExec = st.db.ExecContext(ctx, sqlEmail)
	}
	if errExec != nil {
		return errExec
	}

	return nil
}

// MigrateDown drops the realm store tables
func (st *storeImplementation) MigrateDown(ctx context.Context, tx ...*sql.Tx) error {
	var txToUse *sql.Tx
	if len(tx) > 0 {
		txToUse = tx[0]
	}

	sqlEmail := st.SqlDropEmailTable()

	var errExec error
	if txToUse != nil {
		_, errExec = txToUse.ExecContext(ctx, sqlEmail)
	} else {
		_, errExec = st.db.ExecContext(ctx, sqlEmail)
	}
	if errExec != nil {
		return errExec
	}

	sqlPerson := st.SqlDropPersonTable()

	if txToUse != nil {
		_, errExec = txToUse.ExecContext(ctx, sqlPerson)
	} else {
		_, errExec = st.db.ExecContext(ctx, sqlPerson)
	}
	if errExec != nil {
		return errExec
	}

	return nil
}

func (st *storeImplementation) DriverName(db *sql.DB) string {
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

func (st *storeImplementation) DbExec(sqlStr string) error {
	_, err := st.db.Exec(sqlStr)
	return err
}

func (st *storeImplementation) EnableDebug(debug bool) {
	st.debug = debug
}

func (st *storeImplementation) PersonFindOrCreate(firstName, lastName, email string) (PersonInterface, error) {
	person, err := st.PersonFindByEmail(email)
	if err != nil {
		return nil, err
	}

	if person != nil {
		return person, nil
	}

	person = NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName(firstName)
	person.SetLastName(lastName)

	err = st.PersonCreate(person)
	if err != nil {
		return nil, err
	}

	newEmail := NewEmail()
	newEmail.SetID(uid.MicroUid())
	newEmail.SetIsDefault(Yes)
	newEmail.SetStatus(EmailStatusActive)
	newEmail.SetOwnerID(person.GetID())
	newEmail.SetEmailAddress(email)
	isEmailCreated, err := st.EmailCreate(newEmail)
	if err != nil {
		return nil, err
	}

	if !isEmailCreated {
		return nil, ErrEmailCreationFailed
	}

	return person, nil
}
