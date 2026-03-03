package realmstore

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

func (st *storeImplementation) emailQuery(options *EmailQueryOptions) *goqu.SelectDataset {
	q := goqu.Dialect(st.dbDriverName).From(st.emailTableName)

	if value := options.Search; value != "" {
		q = q.Where(
			goqu.Or(
				goqu.C("email_address").ILike(fmt.Sprintf("%%%v%%", options.Search)),
			))
	}

	if value := options.Status; value != "" {
		q = q.Where(goqu.C("status").Eq(options.Status))
	}

	if value := options.OwnerID; value != "" {
		q = q.Where(goqu.C("owner_id").Eq(options.OwnerID))
	}

	if value := options.CreatedAt; value != "" {
		q = q.Where(goqu.C("created_at").Eq(options.CreatedAt))
	}

	if options.IsDefault == Yes {
		q = q.Where(goqu.C("is_default").Eq(Yes))
	}
	if options.IsDefault == No {
		q = q.Where(goqu.C("is_default").Eq(No))
	}

	if value := options.DeletedOnly; value {
		q = q.Where(goqu.C("deleted_at").IsNotNull())
	} else {
		q = q.Where(goqu.C("deleted_at").IsNull())
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
