package realmstore

import (
	"context"
	"log"

	// "github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/sqlscan"
)

type PersonWithEmail struct {
	Person
	EmailAddress *string `json:"email_address" db:"email_address"`
}

func (st *Store) PersonList(options *PersonQueryOptions) ([]PersonWithEmail, error) {
	var sqlStr string

	q := st.personQuery(options)

	sqlStr, _, _ = q.ToSQL()

	log.Println(sqlStr)

	if st.debug {
		log.Println(sqlStr)
	}

	persons := make([]PersonWithEmail, 0)
	err := sqlscan.Select(context.Background(), st.db, &persons, sqlStr)

	if err != nil {
		if sqlscan.NotFound(err) {
			return nil, err
		}

		log.Println("RealmStore. PersoList. Error: ", err)
		return nil, err
	}

	return persons, nil
}


//Old Version
// PersonList creates a person
// func (st *Store) PersonList(options map[string]interface{}, page uint) ([]PersonWithEmail, error) {
// 	var sqlStr string

// 	q := goqu.Dialect(st.dbDriverName).From(st.personTableName)

// 	if value, ok := options["persons_per_page"]; ok && value != 0 {
// 		q = q.Limit(value.(uint)).Offset(value.(uint) * page)
// 	}

// 	if value, ok := options["first_name"]; ok && value != "" {
// 		q = q.Where(goqu.C("first_name").Eq(options["first_name"].(string)))
// 	}

// 	if value, ok := options["last_name"]; ok && value != "" {
// 		q = q.Where(goqu.C("last_name").Eq(options["last_name"].(string)))
// 	}

// 	sortOrder := "desc"
// 	if value, ok := options["sort_order"]; ok && value != "" {
// 		sortOrder = options["sort_order"].(string)
// 	}

// 	if value, ok := options["sort_by"]; ok && value != "" {
// 		if sortOrder == "asc" {
// 			q = q.Order(goqu.I(options["sort_by"].(string)).Asc())
// 		} else {
// 			q = q.Order(goqu.I(options["sort_by"].(string)).Desc())
// 		}
// 	}

// 	if value, ok := options["deleted_only"]; ok && value != "deleted" {
// 		q = q.Where(goqu.C("deleted_at").Table(st.personTableName).IsNotNull()).Select()
// 	} else {
// 		q = q.Where(goqu.C("deleted_at").Table(st.personTableName).IsNull()).Select()
// 	}

// 	if value, ok := options["with_email"]; ok && value.(string) == "true" {
// 		q = q.Select(goqu.C("*").Table(st.personTableName), goqu.C("email_address").Table(st.emailTableName)).Join(goqu.T(st.emailTableName), goqu.On(goqu.Ex{st.personTableName + ".id": goqu.I(st.emailTableName + ".owner_id")}))
// 	}

// 	sqlStr, _, _ = q.ToSQL()

// 	if st.debug {
// 		log.Println(sqlStr)
// 	}

// 	persons := make([]PersonWithEmail, 0)
// 	err := sqlscan.Select(context.Background(), st.db, &persons, sqlStr)

// 	if err != nil {
// 		if sqlscan.NotFound(err) {
// 			return nil, err
// 		}

// 		log.Println("RealmStore. PersoList. Error: ", err)
// 		return nil, err
// 	}

// 	return persons, nil
// }
