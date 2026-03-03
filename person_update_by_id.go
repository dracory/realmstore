package realmstore

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
)

// PersonUpdateByID creates a person
func (st *storeImplementation) PersonUpdateByID(personID string, fields map[string]interface{}) error {
	fields["updated_at"] = time.Now()

	sqlStr, _, err := goqu.Dialect(st.dbDriverName).Update(st.personTableName).Set(fields).Where(goqu.C("id").Eq(personID)).ToSQL()

	if st.debug && err != nil {
		log.Println(err.Error())
	}

	return st.DbExec(sqlStr)
}
