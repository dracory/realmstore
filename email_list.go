package realmstore

import (
	"context"
	"log"

	// "github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/sqlscan"
)

func (st *storeImplementation) EmailList(options *EmailQueryOptions) ([]Email, error) {
	var sqlStr string

	q := st.emailQuery(options)

	sqlStr, _, _ = q.ToSQL()

	log.Println(sqlStr)

	if st.debug {
		log.Println(sqlStr)
	}

	emails := make([]Email, 0)
	err := sqlscan.Select(context.Background(), st.db, &emails, sqlStr)

	if err != nil {
		if sqlscan.NotFound(err) {
			return nil, err
		}

		if st.debug {
			log.Println("RealmStore. EmailList. Error: ", err)
		}
		return nil, err
	}

	return emails, nil
}
