package realmstore

import (
	"database/sql"
)

type StoreInterface interface {
	AutoMigrate() error
	DriverName(db *sql.DB) string
	DbExec(sqlStr string) error
	EnableDebug(debug bool)

	PersonFindOrCreate(firstName, lastName, email string) (*Person, error)
	PersonCreate(person *Person) error
	PersonCreateWithEmail(fields map[string]string) (*Person, error)
	PersonUpdate(person *Person) error
	PersonUpdateByID(personID string, fields map[string]interface{}) error
	PersonFindByID(personID string) (*Person, error)
	PersonFindByEmail(email string) (*Person, error)
	PersonFindDefaultEmail(person_id string) (*Email, error)
	PersonDelete(personId string) error
	PersonList(options *PersonQueryOptions) ([]PersonWithEmail, error)
	PersonMaxPages(options *PersonQueryOptions) (uint, error)

	EmailCreate(email *Email) (bool, error)
	EmailUpdate(email *Email) bool
	EmailUpdateByID(emailId string, fields map[string]interface{}) bool
	EmailFindByID(id string) (*Email, error)
	EmailFindByEmailAddress(emailAddress string) *Email
	EmailFind(options *EmailQueryOptions) (*Email, error)
	EmailDeleteByID(emailId string) (bool, error)
	EmailList(options *EmailQueryOptions) ([]Email, error)
	EmailFoundByPersonId(personId string) (*Email, error)
}
