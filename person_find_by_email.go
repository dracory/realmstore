package realmstore

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/sqlscan"
)

// PersonFindByEmail finds a person by email
func (st *storeImplementation) PersonFindByEmail(email string) (PersonInterface, error) {
	recordEmail := st.EmailFindByEmailAddress(email)

	if recordEmail == nil {
		return nil, nil
	}

	q := goqu.Dialect(st.dbDriverName).From(st.personTableName)
	q = q.Where(goqu.C("id").Eq(recordEmail.GetOwnerID()))
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
		if sqlscan.NotFound(err) {
			return nil, nil
		}

		log.Println("RealmStore. PersonFindByEmail. Error: ", err)
		return nil, err
	}

	return &record, nil
}
