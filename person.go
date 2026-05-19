package realmstore

import "github.com/dracory/sb"

const (
	// PersonStatusActive status active
	PersonStatusActive = "active"
	// PersonStatusInactive status inactive
	PersonStatusInactive = "inactive"
	//PersonRowsPerPage return how many person will have on the page
	PersonRowsPerPage uint = 20
)

// SqlCreatePersonTable returns a SQL string for creating the setting table
func (st *storeImplementation) SqlCreatePersonTable() string {
	sqlMysql := `
	CREATE TABLE IF NOT EXISTS ` + st.personTableName + ` (
	  id varchar(40) NOT NULL PRIMARY KEY,
	  status varchar(40) NOT NULL,
	  first_name varchar(100) NOT NULL,
	  middle_names varchar(100) NOT NULL,
	  last_name varchar(100) NOT NULL,
	  birth_day date DEFAULT NULL,
	  created_at datetime NOT NULL,
	  updated_at datetime NOT NULL,
	  deleted_at datetime DEFAULT NULL
	);
	`

	sqlPostgres := `
	CREATE TABLE IF NOT EXISTS "` + st.personTableName + `" (
	  "id" varchar(40) NOT NULL PRIMARY KEY,
	  "status" varchar(40) NOT NULL,
	  "first_name" varchar(100) NOT NULL,
	  "middle_names" varchar(100) DEFAULT NULL,
	  "last_name" varchar(100) NOT NULL,
	  "birth_day" date DEFAULT NULL,
	  "created_at" timestamptz(6) NOT NULL,
	  "updated_at" timestamptz(6) NOT NULL,
	  "deleted_at" timestamptz(6) DEFAULT NULL
	)
	`

	sqlSqlite := `
	CREATE TABLE IF NOT EXISTS "` + st.personTableName + `" (
	  "id" varchar(40) NOT NULL PRIMARY KEY,
	  "status" varchar(40) NOT NULL,
	  "first_name" varchar(50) NOT NULL,
	  "middle_names" varchar(50) DEFAULT NULL,
	  "last_name" varchar(50) NOT NULL,
	  "birth_day" date DEFAULT NULL,
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

// SqlDropPersonTable returns a SQL string for dropping the person table
func (st *storeImplementation) SqlDropPersonTable() string {
	sql, _ := sb.NewBuilder(st.dbDriverName).
		Table(st.personTableName).
		Drop()
	return sql
}
