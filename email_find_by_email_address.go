package realmstore

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/sqlscan"
)

// EmailFindByEmailAddress finds a emaii by email address
func (st *storeImplementation) EmailFindByEmailAddress(emailAddress string) *Email {
	sqlStr, _, _ := goqu.Dialect(st.dbDriverName).
		From(st.emailTableName).
		Where(goqu.C("email_address").Eq(emailAddress), goqu.C("deleted_at").IsNull()).
		Select("*").
		Limit(1).
		ToSQL()

	if st.debug {
		log.Println(sqlStr)
	}

	var email Email
	err := sqlscan.Get(context.Background(), st.db, &email, sqlStr)
	if err != nil {
		if sqlscan.NotFound(err) {
			return nil
		}

		log.Println("RealmStore. EmailFindByEmailAddress. Error: ", err)
		return nil
	}

	return &email
}
