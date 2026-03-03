package realmstore

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
	uid "github.com/dracory/uid"
)

// EmailCreate creates an email
func (st *storeImplementation) EmailCreate(email EmailInterface) (bool, error) {
	if email.GetID() == "" {
		time.Sleep(1 * time.Millisecond)
		email.SetID(uid.MicroUid())
	}

	email.SetCreatedAt(time.Now())
	email.SetUpdatedAt(time.Now())

	e := email.(*emailImplementation)
	var sqlStr string
	sqlStr, _, _ = goqu.Dialect(st.dbDriverName).Insert(st.emailTableName).Rows(e).ToSQL()

	if st.debug {
		log.Println(sqlStr)
	}

	_, err := st.db.Exec(sqlStr)
	if err != nil {
		return false, err
	}

	return true, nil
}
