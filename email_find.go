package realmstore

// "github.com/doug-martin/goqu/v9"

// EmailFind finds a single email. For multiple emails use EmailList
func (st *storeImplementation) EmailFind(options *EmailQueryOptions) (EmailInterface, error) {

	options.Limit = 1 // we only look for 1 email

	emails, errList := st.EmailList(options)
	if errList != nil {
		return nil, errList
	}

	if len(emails) > 0 {
		return emails[0], nil
	}

	return nil, nil
}
