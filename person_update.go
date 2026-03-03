package realmstore

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
)

// PersonUpdate creates a person
func (st *storeImplementation) PersonUpdate(person PersonInterface) error {
	person.SetUpdatedAt(time.Now())

	p := person.(*Person)
	var sqlStr string
	sqlStr, _, err := goqu.Dialect(st.dbDriverName).
		Update(st.personTableName).
		Set(p).
		Where(goqu.C("id").Eq(p.ID)).
		ToSQL()

	if st.debug && err != nil {
		log.Println(err.Error())
	}

	return st.DbExec(sqlStr)
}
