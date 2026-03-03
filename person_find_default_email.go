package realmstore

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/sqlscan"
)

func (st *Store) PersonFindDefaultEmail(person_id string) (*Email, error) {
	var email Email

	sqlStr, _, err := goqu.Dialect(st.dbDriverName).From(st.emailTableName).Select(email).Where(goqu.C("owner_id").Eq(person_id), goqu.C("is_default").Eq(Yes), goqu.C("deleted_at").IsNull()).ToSQL()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	errSelect := sqlscan.Get(context.Background(), st.db, &email, sqlStr)
	if errSelect != nil {
		log.Println(errSelect)
		return nil, errSelect
	}

	return &email, nil
}
