package realmstore

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
)

// PersonDelete set delete date to person
func (st *storeImplementation) PersonDelete(personId string) error {
	sqlStr, _, err := goqu.Dialect(st.dbDriverName).Update(st.personTableName).Set(goqu.C("deleted_at").Set(time.Now())).Where(goqu.C("id").Eq(personId), goqu.C("deleted_at").IsNull()).ToSQL()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	errDelete := st.DbExec(sqlStr)

	if errDelete != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
