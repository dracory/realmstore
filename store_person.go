package realmstore

import (
	"context"
	"errors"
	"uuid"

	"github.com/dromara/carbon/v2"
)

// == PERSON CRUD =============================================================

// PersonCreate creates a person
func (st *storeImplementation) PersonCreate(ctx context.Context, person PersonInterface) error {
	if person == nil {
		return errors.New("person is nil")
	}

	if person.GetID() == "" {
		person.SetID(uuid.NewV7().String())
	}

	person.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString())
	person.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	row := map[string]any{
		COLUMN_ID:              person.GetID(),
		COLUMN_STATUS:          person.GetStatus(),
		COLUMN_FIRST_NAME:      person.GetFirstName(),
		COLUMN_MIDDLE_NAMES:    person.GetMiddleNames(),
		COLUMN_LAST_NAME:       person.GetLastName(),
		COLUMN_CREATED_AT:      person.GetCreatedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:      person.GetUpdatedAtCarbon().StdTime(),
		COLUMN_SOFT_DELETED_AT: carbon.Parse(MAX_DATETIME, carbon.UTC).StdTime(),
	}

	if person.GetBirthDay() != nil {
		row[COLUMN_BIRTH_DAY] = *person.GetBirthDay()
	}

	return st.db.Query().Table(st.personTableName).Create(row)
}

// PersonCreateWithEmail creates a person with email.
func (st *storeImplementation) PersonCreateWithEmail(ctx context.Context, fields map[string]string) (PersonInterface, error) {
	person := NewPerson()
	if id, exists := fields["id"]; exists {
		person.SetID(id)
	}

	person.SetFirstName(fields["firstName"])
	person.SetLastName(fields["lastName"])
	person.SetStatus(PersonStatusActive)

	err := st.PersonCreate(ctx, person)
	if err != nil {
		return nil, err
	}

	if emailAddress, exists := fields["email"]; exists {
		email := NewEmail()
		email.SetEmailAddress(emailAddress)
		email.SetStatus(EmailStatusActive)
		email.SetOwnerID(person.GetID())
		err := st.EmailCreate(ctx, email)
		if err != nil {
			return nil, err
		}
	}

	return person, nil
}

// PersonFindOrCreate finds a person by email or creates one if not found.
func (st *storeImplementation) PersonFindOrCreate(ctx context.Context, firstName, lastName, email string) (PersonInterface, error) {
	person, err := st.PersonFindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if person != nil {
		return person, nil
	}

	person = NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName(firstName)
	person.SetLastName(lastName)

	err = st.PersonCreate(ctx, person)
	if err != nil {
		return nil, err
	}

	newEmail := NewEmail()
	newEmail.SetStatus(EmailStatusActive)
	newEmail.SetIsDefault(Yes)
	newEmail.SetOwnerID(person.GetID())
	newEmail.SetEmailAddress(email)
	err = st.EmailCreate(ctx, newEmail)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// PersonFindByID finds a person by ID
func (st *storeImplementation) PersonFindByID(ctx context.Context, personID string) (PersonInterface, error) {
	if personID == "" {
		return nil, errors.New("person id is empty")
	}

	list, err := st.PersonList(ctx, PersonQuery().
		SetID(personID).
		SetLimit(1))
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

// PersonFindByEmail finds a person by email
func (st *storeImplementation) PersonFindByEmail(ctx context.Context, email string) (PersonInterface, error) {
	recordEmail, err := st.EmailFindByEmailAddress(ctx, email)
	if err != nil {
		return nil, err
	}

	if recordEmail == nil {
		return nil, nil
	}

	return st.PersonFindByID(ctx, recordEmail.GetOwnerID())
}

// PersonFindDefaultEmail finds the default email for a person
func (st *storeImplementation) PersonFindDefaultEmail(ctx context.Context, personID string) (EmailInterface, error) {
	if personID == "" {
		return nil, errors.New("person id is empty")
	}

	return st.EmailFind(ctx, EmailQuery().
		SetOwnerID(personID).
		SetIsDefault(Yes).
		SetLimit(1))
}

// PersonDelete soft-deletes a person by ID
func (st *storeImplementation) PersonDelete(ctx context.Context, personID string) error {
	if personID == "" {
		return errors.New("person id is empty")
	}

	row := map[string]any{
		COLUMN_SOFT_DELETED_AT: carbon.Now(carbon.UTC).StdTime(),
		COLUMN_UPDATED_AT:      carbon.Now(carbon.UTC).StdTime(),
	}

	_, err := st.db.Query().
		Table(st.personTableName).
		Where(COLUMN_ID+" = ?", personID).
		Update(row)

	return err
}

// PersonList lists persons based on a query
func (st *storeImplementation) PersonList(ctx context.Context, query PersonQueryInterface) ([]PersonWithEmailInterface, error) {
	if query == nil {
		query = PersonQuery()
	}

	if err := query.Validate(); err != nil {
		return nil, err
	}

	q := st.buildPersonQuery(query)

	var results []map[string]any
	err := q.Get(&results)
	if err != nil {
		return nil, err
	}

	list := make([]PersonWithEmailInterface, 0, len(results))
	for _, result := range results {
		model := st.mapToPersonWithEmail(result)
		list = append(list, model)
	}

	return list, nil
}

// PersonCount returns the total number of persons matching the query filters
func (st *storeImplementation) PersonCount(ctx context.Context, query PersonQueryInterface) (int64, error) {
	if query == nil {
		query = PersonQuery()
	}

	if err := query.Validate(); err != nil {
		return 0, err
	}

	q := st.buildPersonQuery(query)

	var count int64
	err := q.Count(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// PersonMaxPages return max page based on limit
func (st *storeImplementation) PersonMaxPages(ctx context.Context, query PersonQueryInterface) (uint, error) {
	count, err := st.PersonCount(ctx, query)
	if err != nil {
		return 0, err
	}

	limit := uint(PersonRowsPerPage)
	if query != nil && query.IsLimitSet() && query.GetLimit() > 0 {
		limit = uint(query.GetLimit())
	}

	if limit == 0 {
		return 0, nil
	}

	pageCount := uint(count) / limit
	if uint(count)%limit > 0 {
		pageCount++
	}

	return pageCount, nil
}

// PersonUpdate updates a person
func (st *storeImplementation) PersonUpdate(ctx context.Context, person PersonInterface) error {
	if person == nil {
		return errors.New("person is nil")
	}

	person.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	row := map[string]any{
		COLUMN_STATUS:       person.GetStatus(),
		COLUMN_FIRST_NAME:   person.GetFirstName(),
		COLUMN_MIDDLE_NAMES: person.GetMiddleNames(),
		COLUMN_LAST_NAME:    person.GetLastName(),
		COLUMN_UPDATED_AT:   person.GetUpdatedAtCarbon().StdTime(),
	}

	if person.GetBirthDay() != nil {
		row[COLUMN_BIRTH_DAY] = *person.GetBirthDay()
	}

	_, err := st.db.Query().
		Table(st.personTableName).
		Where(COLUMN_ID+" = ?", person.GetID()).
		Update(row)

	return err
}

// PersonUpdateByID updates a person by ID with partial fields
func (st *storeImplementation) PersonUpdateByID(ctx context.Context, personID string, fields map[string]any) error {
	if personID == "" {
		return errors.New("person id is empty")
	}

	fields[COLUMN_UPDATED_AT] = carbon.Now(carbon.UTC).StdTime()

	_, err := st.db.Query().
		Table(st.personTableName).
		Where(COLUMN_ID+" = ?", personID).
		Update(fields)

	return err
}

// == HELPERS =================================================================

func (st *storeImplementation) mapToPersonWithEmail(data map[string]any) *personWithEmail {
	p := &personWithEmail{}

	if v, ok := data[COLUMN_ID].(string); ok {
		p.SetID(v)
	}
	if v, ok := data[COLUMN_STATUS].(string); ok {
		p.SetStatus(v)
	}
	if v, ok := data[COLUMN_FIRST_NAME].(string); ok {
		p.SetFirstName(v)
	}
	if v, ok := data[COLUMN_MIDDLE_NAMES].(string); ok {
		p.SetMiddleNames(v)
	}
	if v, ok := data[COLUMN_LAST_NAME].(string); ok {
		p.SetLastName(v)
	}
	if v, ok := data[COLUMN_CREATED_AT]; ok {
		switch vt := v.(type) {
		case string:
			p.SetCreatedAt(vt)
		case []byte:
			p.SetCreatedAt(string(vt))
		}
	}
	if v, ok := data[COLUMN_UPDATED_AT]; ok {
		switch vt := v.(type) {
		case string:
			p.SetUpdatedAt(vt)
		case []byte:
			p.SetUpdatedAt(string(vt))
		}
	}
	if v, ok := data[COLUMN_SOFT_DELETED_AT]; ok {
		switch vt := v.(type) {
		case string:
			p.SetSoftDeletedAt(vt)
		case []byte:
			p.SetSoftDeletedAt(string(vt))
		}
	}
	if v, ok := data[COLUMN_EMAIL_ADDRESS].(string); ok {
		p.SetEmailAddress(v)
	}

	return p
}
