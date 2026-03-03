package realmstore

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

func (st *Store) personQuery(options *PersonQueryOptions) *goqu.SelectDataset {
	q := goqu.Dialect(st.dbDriverName).From(st.personTableName)

	if options.WithEmail {
		q = q.Select(goqu.C("*").Table(st.personTableName), goqu.C("email_address").Table(st.emailTableName)).Join(goqu.T(st.emailTableName), goqu.On(goqu.Ex{st.personTableName + ".id": goqu.I(st.emailTableName + ".owner_id")}))
	}

	if value := options.Search; value != "" {
		q = q.Where(
			goqu.Or(
				goqu.C("first_name").ILike(fmt.Sprintf("%%%v%%",options.Search)),
				goqu.C("last_name").ILike(fmt.Sprintf("%%%v%%",options.Search)),
				goqu.C("email_address").ILike(fmt.Sprintf("%%%v%%",options.Search)),
			))
	}

	// if value := options.Last_name; value != "" {
	// 	q = q.Where(goqu.C("last_name").Eq(options.Last_name))
	// }

	if value := options.Created_at; value != "" {
		q = q.Where(goqu.C("created_at").Eq(options.Created_at))
	}

	if value := options.DeletedOnly; value {
		q = q.Where(goqu.C("deleted_at").Table(st.personTableName).IsNotNull())
	} else {
		q = q.Where(goqu.C("deleted_at").Table(st.personTableName).IsNull())
	}

	if !options.CountOnly {
		if options.Limit > 0 {
			q = q.Limit(uint(options.Limit))
		}

		if options.Offset > 0 {
			q = q.Offset(uint(options.Offset))
		}
	}

	sortOrder := "desc"
	if options.SortOrder != "" {
		sortOrder = options.SortOrder
	}

	if options.SortBy != "" {
		if sortOrder == "asc" {
			q = q.Order(goqu.I(options.SortBy).Asc())
		} else {
			q = q.Order(goqu.I(options.SortBy).Desc())
		}
	}

	return q
}
