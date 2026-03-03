package realmstore

import (
	"context"
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/sqlscan"
)

// PersonFindByID finds a person by ID
func (st *storeImplementation) PersonFindByID(personID string) (PersonInterface, error) {
	q := goqu.Dialect(st.dbDriverName).From(st.personTableName)
	q = q.Where(goqu.C("id").Eq(personID))
	sqlStr, _, sqlErr := q.Where(goqu.C("deleted_at").IsNull()).Select().Limit(1).ToSQL()

	if st.debug {
		log.Println(sqlStr)
	}

	if sqlErr != nil {
		return nil, sqlErr
	}

	var record personImplementation
	err := sqlscan.Get(context.Background(), st.db, &record, sqlStr)

	if err != nil {
		if err == sql.ErrNoRows {
			// sqlscan does not use this anymore
			return nil, nil
		}

		if sqlscan.NotFound(err) {
			return nil, nil
		}

		log.Fatal("RealmSTore. PersonFindByID. Error: ", err)
		return nil, err
	}

	return &record, nil
}
