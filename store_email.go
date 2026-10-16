package realmstore

import (
	"context"
	"errors"
	"uuid"

	"github.com/dromara/carbon/v2"
)

// == EMAIL CRUD ==============================================================

// EmailCreate creates an email
func (st *storeImplementation) EmailCreate(ctx context.Context, email EmailInterface) error {
	if email == nil {
		return errors.New("email is nil")
	}

	if email.GetID() == "" {
		email.SetID(uuid.NewV7().String())
	}

	email.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString())
	email.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	row := map[string]any{
		COLUMN_ID:              email.GetID(),
		COLUMN_OWNER_ID:        email.GetOwnerID(),
		COLUMN_STATUS:          email.GetStatus(),
		COLUMN_EMAIL_ADDRESS:   email.GetEmailAddress(),
		COLUMN_IS_DEFAULT:      email.GetIsDefault(),
		COLUMN_CREATED_AT:      email.GetCreatedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:      email.GetUpdatedAtCarbon().StdTime(),
		COLUMN_SOFT_DELETED_AT: carbon.Parse(MAX_DATETIME, carbon.UTC).StdTime(),
	}

	return st.db.Query().Table(st.emailTableName).Create(row)
}

// EmailDeleteByID soft-deletes an email by ID
func (st *storeImplementation) EmailDeleteByID(ctx context.Context, emailID string) error {
	if emailID == "" {
		return errors.New("email id is empty")
	}

	row := map[string]any{
		COLUMN_SOFT_DELETED_AT: carbon.Now(carbon.UTC).StdTime(),
		COLUMN_UPDATED_AT:      carbon.Now(carbon.UTC).StdTime(),
	}

	_, err := st.db.Query().
		Table(st.emailTableName).
		Where(COLUMN_ID+" = ?", emailID).
		Update(row)

	return err
}

// EmailFind finds a single email. For multiple emails use EmailList
func (st *storeImplementation) EmailFind(ctx context.Context, query EmailQueryInterface) (EmailInterface, error) {
	if query == nil {
		query = EmailQuery()
	}

	query.SetLimit(1)

	list, err := st.EmailList(ctx, query)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

// EmailFindByID finds an email by ID
func (st *storeImplementation) EmailFindByID(ctx context.Context, id string) (EmailInterface, error) {
	if id == "" {
		return nil, errors.New("email id is empty")
	}

	return st.EmailFind(ctx, EmailQuery().
		SetID(id).
		SetLimit(1))
}

// EmailFindByEmailAddress finds an email by email address
func (st *storeImplementation) EmailFindByEmailAddress(ctx context.Context, emailAddress string) (EmailInterface, error) {
	if emailAddress == "" {
		return nil, errors.New("email address is empty")
	}

	return st.EmailFind(ctx, EmailQuery().
		SetEmailAddress(emailAddress).
		SetLimit(1))
}

// EmailFindByPersonID finds an email by person ID
func (st *storeImplementation) EmailFindByPersonID(ctx context.Context, personID string) (EmailInterface, error) {
	if personID == "" {
		return nil, errors.New("person id is empty")
	}

	return st.EmailFind(ctx, EmailQuery().
		SetOwnerID(personID).
		SetLimit(1))
}

// EmailList lists emails based on a query
func (st *storeImplementation) EmailList(ctx context.Context, query EmailQueryInterface) ([]EmailInterface, error) {
	if query == nil {
		query = EmailQuery()
	}

	if err := query.Validate(); err != nil {
		return nil, err
	}

	q := st.buildEmailQuery(query)

	var results []map[string]any
	err := q.Get(&results)
	if err != nil {
		return nil, err
	}

	list := make([]EmailInterface, 0, len(results))
	for _, result := range results {
		model := st.mapToEmail(result)
		list = append(list, model)
	}

	return list, nil
}

// EmailCount returns the total number of emails matching the query filters
func (st *storeImplementation) EmailCount(ctx context.Context, query EmailQueryInterface) (int64, error) {
	if query == nil {
		query = EmailQuery()
	}

	if err := query.Validate(); err != nil {
		return 0, err
	}

	q := st.buildEmailQuery(query)

	var count int64
	err := q.Count(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// EmailUpdate updates an email
func (st *storeImplementation) EmailUpdate(ctx context.Context, email EmailInterface) error {
	if email == nil {
		return errors.New("email is nil")
	}

	email.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	row := map[string]any{
		COLUMN_OWNER_ID:      email.GetOwnerID(),
		COLUMN_STATUS:        email.GetStatus(),
		COLUMN_EMAIL_ADDRESS: email.GetEmailAddress(),
		COLUMN_IS_DEFAULT:    email.GetIsDefault(),
		COLUMN_UPDATED_AT:    email.GetUpdatedAtCarbon().StdTime(),
	}

	_, err := st.db.Query().
		Table(st.emailTableName).
		Where(COLUMN_ID+" = ?", email.GetID()).
		Update(row)

	return err
}

// EmailUpdateByID updates an email by ID with partial fields
func (st *storeImplementation) EmailUpdateByID(ctx context.Context, emailID string, fields map[string]any) error {
	if emailID == "" {
		return errors.New("email id is empty")
	}

	fields[COLUMN_UPDATED_AT] = carbon.Now(carbon.UTC).StdTime()

	_, err := st.db.Query().
		Table(st.emailTableName).
		Where(COLUMN_ID+" = ?", emailID).
		Update(fields)

	return err
}

// == HELPERS =================================================================

func (st *storeImplementation) mapToEmail(data map[string]any) *emailImplementation {
	e := &emailImplementation{}

	if v, ok := data[COLUMN_ID].(string); ok {
		e.SetID(v)
	}
	if v, ok := data[COLUMN_OWNER_ID].(string); ok {
		e.SetOwnerID(v)
	}
	if v, ok := data[COLUMN_STATUS].(string); ok {
		e.SetStatus(v)
	}
	if v, ok := data[COLUMN_EMAIL_ADDRESS].(string); ok {
		e.SetEmailAddress(v)
	}
	if v, ok := data[COLUMN_IS_DEFAULT].(string); ok {
		e.SetIsDefault(v)
	}
	if v, ok := data[COLUMN_CREATED_AT]; ok {
		switch vt := v.(type) {
		case string:
			e.SetCreatedAt(vt)
		case []byte:
			e.SetCreatedAt(string(vt))
		}
	}
	if v, ok := data[COLUMN_UPDATED_AT]; ok {
		switch vt := v.(type) {
		case string:
			e.SetUpdatedAt(vt)
		case []byte:
			e.SetUpdatedAt(string(vt))
		}
	}
	if v, ok := data[COLUMN_SOFT_DELETED_AT]; ok {
		switch vt := v.(type) {
		case string:
			e.SetSoftDeletedAt(vt)
		case []byte:
			e.SetSoftDeletedAt(string(vt))
		}
	}

	return e
}
