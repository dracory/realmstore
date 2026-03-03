package realmstore

import (
	"context"
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/sqlscan"
)

const (
	// EmailStatusActive actove status
	EmailStatusActive = "active"
	// EmailStatusInactive inactive status
	EmailStatusInactive = "inactive"
	// EmailStatusPendingActivation pending_activation status
	EmailStatusPendingActivation = "pending_activation"

	// EmailTypePersonal represents a personal type
	EmailTypePersonal = "personal"
	// EmailTypeWork represents a work type
	EmailTypeWork = "work"

	// No represents a boolean no
	No = "no"

	// Yes represents a boolean yes
	Yes = "yes"
)

// SqlCreateEmailTable returns a SQL string for creating the user table
func (st *storeImplementation) SqlCreateEmailTable() string {
	sqlMysql := `
	CREATE TABLE IF NOT EXISTS ` + st.emailTableName + ` (
	  id varchar(40) NOT NULL PRIMARY KEY,
	  status varchar(40) NOT NULL,
	  owner_id varchar(40) NOT NULL, 
	  email_address varchar(100) NOT NULL,
	  is_default varchar(3) NOT NULL,
	  created_at datetime NOT NULL,
	  updated_at datetime NOT NULL,
	  deleted_at datetime DEFAULT NULL
	);
	`

	sqlPostgres := `
	CREATE TABLE IF NOT EXISTS "` + st.emailTableName + `" (
	  "id" varchar(40) NOT NULL PRIMARY KEY,
	  "status" varchar(40) NOT NULL,
	  "owner_id" varchar(40) NOT NULL,
	  "email_address" varchar(100) NOT NULL,
	  "is_default" varchar(3) NOT NULL,
	  "created_at" datetime NOT NULL,
	  "updated_at" datetime NOT NULL,
	  "deleted_at" datetime DEFAULT NULL
	)
	`

	sqlSqlite := `
	CREATE TABLE IF NOT EXISTS "` + st.emailTableName + `" (
	  "id" varchar(40) NOT NULL PRIMARY KEY,
	  "status" varchar(40) NOT NULL,
	  "owner_id" varchar(40) NOT NULL,
	  "email_address" varchar(100) NOT NULL,
	  "is_default" varchar(3) NOT NULL,
	  "created_at" datetime NOT NULL,
	  "updated_at" datetime NOT NULL,
	  "deleted_at" datetime DEFAULT NULL
	)
	`

	sql := "unsupported driver " + st.dbDriverName

	if st.dbDriverName == "mysql" {
		sql = sqlMysql
	}
	if st.dbDriverName == "postgres" {
		sql = sqlPostgres
	}
	if st.dbDriverName == "sqlite" {
		sql = sqlSqlite
	}

	return sql
}

func (st *storeImplementation) EmailFindByPersonId(personId string) (EmailInterface, error) {
	sqlStr, _, sqlErr := goqu.Dialect(st.dbDriverName).From(st.emailTableName).Where(goqu.C("owner_id").Eq(personId), goqu.C("deleted_at").IsNull()).Limit(1).ToSQL()

	if sqlErr != nil {
		return nil, sqlErr
	}

	var record emailImplementation
	err := sqlscan.Get(context.Background(), st.db, &record, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		if sqlscan.NotFound(err) {
			return nil, nil
		}

		log.Println("FunnelStore. EmailFindByPersonId. Error: ", err)
		return nil, err
	}
	return &record, nil
}
