package realmstore

import (
	"time"
)

const (
	// PersonStatusActive status active
	PersonStatusActive = "active"
	// PersonStatusInactive status inactive
	PersonStatusInactive = "inactive"
	//PersonRowsPerPage return how many person will have on the page
	PersonRowsPerPage uint = 20
)

// Person type
type Person struct {
	ID          string `json:"id" db:"id"`                     // varchar 40, primary key
	Status      string `json:"status" db:"status"`             // varchar 40
	FirstName   string `json:"first_name" db:"first_name"`     // varchar 100
	MiddleNames string `json:"middle_names" db:"middle_names"` // varchar 100
	LastName    string `json:"last_name" db:"last_name"`       // varchar 100
	// Email       string     `json:"email" db:"email"`                     // varchar 255
	BirthDay  *time.Time `json:"birth_day" db:"birth_day"`             // date
	CreatedAt time.Time  `json:"created_at" db:"created_at"`           // datetime not null
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`           // datetime not null
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"` // datetime
}

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
