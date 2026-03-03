package realmstore

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	uid "github.com/dracory/uid"
)

// PersonCreateWithEmail creates a person with email.
func (st *storeImplementation) PersonCreateWithEmail(fields map[string]string) (*Person, error) {
	person := Person{}
	if _, exists := fields["id"]; !exists {
		time.Sleep(1 * time.Millisecond) // !!! important
		person.ID = uid.MicroUid()
	}

	person.FirstName = fields["firstName"]
	person.LastName = fields["lastName"]

	person.CreatedAt = time.Now()
	person.UpdatedAt = time.Now()
	person.Status = PersonStatusActive

	var sqlStr string
	sqlStr, _, _ = goqu.Dialect(st.dbDriverName).Insert(st.personTableName).Rows(person).ToSQL()

	err := st.DbExec(sqlStr)

	if err != nil {
		return nil, err
	}

	if emailAddress, exists := fields["email"]; exists {
		isOk, errEmail := st.EmailCreate(&Email{EmailAddress: emailAddress, Status: "active", OwnerID: person.ID})
		if errEmail != nil {
			return nil, errEmail
		}
		if !isOk {
			return nil, errEmail
		}
	}

	return &person, nil
}
