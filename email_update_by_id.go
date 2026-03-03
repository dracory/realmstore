package realmstore

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
)

func (st *storeImplementation) EmailUpdateByID(emailId string, filds map[string]interface{}) bool {
	filds["updated_at"] = time.Now()

	sqlStr, _, err := goqu.Dialect(st.dbDriverName).Update(st.emailTableName).Set(filds).Where(goqu.C("id").Eq(emailId)).ToSQL()
	if err != nil {
		log.Println("Error generating SQL:", err.Error())
		return false
	}

	err = st.DbExec(sqlStr)
	if err != nil {
		log.Println("Error executing SQL:", err.Error())
		return false
	}

	return true
}
