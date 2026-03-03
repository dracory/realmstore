package realmstore

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
)

// PersonUpdate creates a person
func (st *storeImplementation) PersonUpdate(person *Person) error {
	person.UpdatedAt = time.Now()

	var sqlStr string
	sqlStr, _, err := goqu.Dialect(st.dbDriverName).
		Update(st.personTableName).
		Set(person).
		Where(goqu.C("id").Eq(person.ID)).
		ToSQL()

	if st.debug && err != nil {
		log.Println(err.Error())
	}

	return st.DbExec(sqlStr)
}
