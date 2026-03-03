package realmstore

import (
	"errors"
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
)

func (st *storeImplementation) EmailDeleteByID(emailId string) (bool, error) {
	sqlStr, _, err := goqu.Dialect(st.dbDriverName).
		Update(st.emailTableName).
		Where(goqu.C("id").Eq(emailId), goqu.C("deleted_at").IsNull()).
		Set(goqu.C("deleted_at").Set(time.Now())).
		ToSQL()

	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	result, errExec := st.db.Exec(sqlStr)
	if errExec != nil {
		return false, errExec
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	if affectedRows < 1 {
		log.Println("Email was not deleted")
		return false, errors.New("Email was not delete")
	}

	return true, nil
}
