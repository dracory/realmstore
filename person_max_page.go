package realmstore

import (
	"log"

	"github.com/doug-martin/goqu/v9"
)

// PersonMaxPages return max page not need CountOnly But need Limit
func (st *Store) PersonMaxPages(options *PersonQueryOptions) (uint, error) {
	q := st.personQuery(options).Select(goqu.COUNT("*"))

	sqlStr, _, errStr := q.ToSQL()

	if errStr != nil {
		log.Println(errStr.Error())
		return 0, errStr
	}

	count := uint(0)
	err := st.db.QueryRow(sqlStr).Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	pageCount := uint(0)
	if value := options.Limit; value != 0 {
		pageCount = count / options.Limit

		if count == options.Limit {
			return pageCount, nil
		}

		return pageCount + 1, nil
	}

	pageCount = count / PersonRowsPerPage
	log.Println(sqlStr, pageCount)

	if count == PersonRowsPerPage {
		return pageCount, nil

	}
	return pageCount + 1, nil
}

// PersonMaxPages return max page Old Version
// func (st *Store) PersonMaxPages(options map[string]interface{}) (uint, error) {
// 	q := goqu.Dialect(st.dbDriverName).From(st.personTableName).Select(goqu.COUNT("*"), goqu.COUNT("deleted_at"))

// 	if value, ok := options["first_name"]; ok && value != "" {
// 		q = q.Where(goqu.C("first_name").Eq(options["first_name"].(string)))
// 	}

// 	if value, ok := options["last_name"]; ok && value != "" {
// 		q = q.Where(goqu.C("last_name").Eq(options["last_name"].(string)))
// 	}

// 	sqlStr, _, errStr := q.ToSQL()

// 	if errStr != nil {
// 		log.Println(errStr.Error())
// 		return 0, errStr
// 	}

// 	countAll := uint(0)
// 	countDeleted := uint(0)
// 	err := st.db.QueryRow(sqlStr).Scan(&countAll, &countDeleted)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return 0, err
// 	}

// 	pageCount := uint(0)
// 	if value, ok := options["persons_per_page"]; ok && value.(uint) != 0 {
// 		pageCount = (countAll - countDeleted) / value.(uint)
// 		return pageCount, nil
// 	}

// 	pageCount = (countAll - countDeleted) / PersonRowsPerPage
// 	return pageCount, nil
// }
