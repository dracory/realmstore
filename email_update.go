package realmstore

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
)

// EmailUpdate update exist email
func (st *storeImplementation) EmailUpdate(email EmailInterface) bool {
	email.SetUpdatedAt(time.Now())

	e := email.(*emailImplementation)
	sqlStr, _, err := goqu.Dialect(st.dbDriverName).Update(st.emailTableName).Set(e).Where(goqu.C("id").Eq(e.ID)).ToSQL()
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
