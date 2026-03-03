package realmstore

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	uid "github.com/dracory/uid"
)

// PersonCreateWithEmail creates a person with email.
func (st *storeImplementation) PersonCreateWithEmail(fields map[string]string) (PersonInterface, error) {
	person := NewPerson()
	if _, exists := fields["id"]; !exists {
		time.Sleep(1 * time.Millisecond) // !!! important
		person.SetID(uid.MicroUid())
	}

	person.SetFirstName(fields["firstName"])
	person.SetLastName(fields["lastName"])

	person.SetCreatedAt(time.Now())
	person.SetUpdatedAt(time.Now())
	person.SetStatus(PersonStatusActive)

	p := person.(*personImplementation)
	var sqlStr string
	sqlStr, _, _ = goqu.Dialect(st.dbDriverName).Insert(st.personTableName).Rows(p).ToSQL()

	err := st.DbExec(sqlStr)

	if err != nil {
		return nil, err
	}

	if emailAddress, exists := fields["email"]; exists {
		isOk, errEmail := st.EmailCreate(&Email{EmailAddress: emailAddress, Status: "active", OwnerID: person.GetID()})
		if errEmail != nil {
			return nil, errEmail
		}
		if !isOk {
			return nil, errEmail
		}
	}

	return person, nil
}
