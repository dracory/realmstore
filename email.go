package realmstore

import (
	"context"
	"database/sql"
	"log"
	"time"

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

// Email type
type Email struct {
	ID           string `json:"id"  db:"id" gorm:"type:varchar(40);column:id;primary_key;"`
	OwnerID      string `json:"owner_id" db:"owner_id" gorm:"type:varchar(40);column:owner_id;"`
	Status       string `json:"status" db:"status" gorm:"type:varchar(40);column:status;"`
	EmailAddress string `json:"email_address" db:"email_address" gorm:"type:varchar(100);column:email_address;unique_index"`
	IsDefault    string `json:"is_default" db:"is_default" gorm:"type:varchar(3);column:is_default;"`
	// Type           string     `json:"type" db:"type" gorm:"type:varchar(255);column:type;DEFAULT NULL;"`
	// IsVerified     string     `json:"is_verified" db:"is_verified" gorm:"type:varchar(10);column:is_verified;"`
	// IsPrimary      string     `json:"is_primary" db:"is_primary" gorm:"type:varchar(10);column:is_primary;DEFAULT NULL;"`
	// LastVerifiedAt time.Time  `json:"last_verified_at" db:"last_verified_at" gorm:"type:datetime;column:last_verified_at;DEFAULT NULL;"`
	CreatedAt time.Time  `json:"created_at" db:"created_at" gorm:"type:datetime;column:created_at;"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at" gorm:"type:datetime;column:updated_at;"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at" gorm:"type:datetime;olumn:deleted_at;DEFAULT NULL;"`
}

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

// EmailFoundByPersonId find email by person id return one row Email struct
func (st *storeImplementation) EmailFoundByPersonId(personId string) (*Email, error) {
	sqlStr, _, sqlErr := goqu.Dialect(st.dbDriverName).From(st.emailTableName).Where(goqu.C("owner_id").Eq(personId), goqu.C("deleted_at").IsNull()).Limit(1).ToSQL()

	if sqlErr != nil {
		return nil, sqlErr
	}

	var record Email
	err := sqlscan.Get(context.Background(), st.db, &record, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			// sqlscan does not use this anymore
			return nil, nil
		}

		if sqlscan.NotFound(err) {
			return nil, nil
		}

		log.Println("FunnelStore. EmailFoundByPersonId. Error: ", err)
		return nil, err
	}
	return &record, nil
}
