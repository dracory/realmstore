package realmstore

import (
	"database/sql"
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/sqlscan"
)

// EmailFindByID finds a email by ID
func (st *storeImplementation) EmailFindByID(id string) (*Email, error) {
	email := &Email{}
	sqlStr, _, _ := goqu.Dialect(st.dbDriverName).Select(&Email{}).From(st.emailTableName).Where(goqu.C("id").Eq(id), goqu.C("deleted_at").IsNull()).ToSQL()

	// if st.debug {
	log.Println(sqlStr)
	// }

	var createdAt string
	var updatedAt string
	var deletedAt *string

	err := st.db.QueryRow(sqlStr).Scan(&createdAt, &deletedAt, &email.EmailAddress, &email.ID, &email.IsDefault, &email.OwnerID, &email.Status, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// sqlscan does not use this anymore
			return nil, nil
		}

		if sqlscan.NotFound(err) {
			return nil, nil
		}

		log.Println("RealmStore. EmailFindByID. Error: ", err)
		return nil, err
	}

	layout := "Mon Jan 02 2006 15:04:05 GMT-0700"
	createdAtTime, err := time.Parse(layout, createdAt)
	if err == nil {
		email.CreatedAt = createdAtTime
	}
	updatedAtTime, err := time.Parse(layout, updatedAt)
	if err == nil {
		email.UpdatedAt = updatedAtTime
	}
	if deletedAt != nil {
		deletedAtTime, err := time.Parse(layout, *deletedAt)
		if err == nil {
			email.DeletedAt = &deletedAtTime
		}
	}

	return email, nil
}
