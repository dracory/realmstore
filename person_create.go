package realmstore

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	uid "github.com/dracory/uid"
)

// PersonCreate creates a person
func (st *storeImplementation) PersonCreate(person PersonInterface) error {
	if person.GetID() == "" {
		time.Sleep(1 * time.Millisecond) // !!! important
		person.SetID(uid.MicroUid())
	}
	person.SetCreatedAt(time.Now())
	person.SetUpdatedAt(time.Now())

	p := person.(*personImplementation)
	sqlStr, _, _ := goqu.Dialect(st.dbDriverName).Insert(st.personTableName).Rows(p).ToSQL()

	return st.DbExec(sqlStr)
}
