package realmstore

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	uid "github.com/dracory/uid"
)

// PersonCreate creates a person
func (st *storeImplementation) PersonCreate(person *Person) error {
	if person.ID == "" {
		time.Sleep(1 * time.Millisecond) // !!! important
		person.ID = uid.MicroUid()
	}
	person.CreatedAt = time.Now()
	person.UpdatedAt = time.Now()

	sqlStr, _, _ := goqu.Dialect(st.dbDriverName).Insert(st.personTableName).Rows(person).ToSQL()

	return st.DbExec(sqlStr)
}
