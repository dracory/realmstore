package realmstore

// type EmailsGetByPersonIDOptions struct {
// 	Status string
// 	// Created_at  string
// 	// SortBy      string
// 	// SortOrder   string // asc / dec
// 	// Limit       uint
// 	// Offset      uint
// 	// CountOnly   bool // returns only the count if enabled
// 	DeletedOnly bool // returns only deleted Emails
// }

// func (st *Store) EmailsGetByPersonID(owner_id string, options EmailsGetByPersonIDOptions) ([]Email, error) {
// 	emails := []Email{}

// 	query := goqu.Dialect(st.dbDriverName).From(st.emailTableName).Select().Where(goqu.C("owner_id").Eq(owner_id))

// 	if options.Status != "" {
// 		query = query.Where(goqu.C("status").Eq(options.Status))
// 	}

// 	if value := options.DeletedOnly; value {
// 		query = query.Where(goqu.C("deleted_at").IsNotNull())
// 	} else {
// 		query = query.Where(goqu.C("deleted_at").IsNull())
// 	}

// 	sqlStr, _, errEmails := query.ToSQL()

// 	if errEmails != nil {
// 		return nil, errEmails
// 	}

// 	err := sqlscan.Select(context.Background(), st.db, &emails, sqlStr)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			// sqlscan does not use this anymore
// 			return nil, err
// 		}

// 		if sqlscan.NotFound(err) {
// 			return nil, err
// 		}

// 		log.Println("RealmStore. EmailsGetByPersonID. Error: ", err)
// 		return nil, err
// 	}

// 	return emails, nil
// }
