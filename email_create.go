package realmstore

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
	uid "github.com/dracory/uid"
)

// EmailCreate creates an email
func (st *Store) EmailCreate(email *Email) (bool, error) {
	if email.ID == "" {
		time.Sleep(1 * time.Millisecond)
		email.ID = uid.MicroUid()
	}

	email.CreatedAt = time.Now()
	email.UpdatedAt = time.Now()

	var sqlStr string
	sqlStr, _, _ = goqu.Dialect(st.dbDriverName).Insert(st.emailTableName).Rows(email).ToSQL()

	if st.debug {
		log.Println(sqlStr)
	}

	_, err := st.db.Exec(sqlStr)
	if err != nil {
		return false, err
	}

	return true, nil
}
