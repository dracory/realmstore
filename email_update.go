package realmstore

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
)

// EmailUpdate update exist email
func (st *Store) EmailUpdate(email *Email) bool {
	email.UpdatedAt = time.Now()

	sqlStr, _, err := goqu.Dialect(st.dbDriverName).Update(st.emailTableName).Set(email).Where(goqu.C("id").Eq(email.ID)).ToSQL()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	errExec := st.DbExec(sqlStr)
	if errExec != nil {
		log.Println(errExec.Error())
		return false
	}

	return true
}
