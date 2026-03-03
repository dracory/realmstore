package realmstore

import (
	"database/sql"
)

type StoreInterface interface {
	AutoMigrate() error
	DriverName(db *sql.DB) string
	DbExec(sqlStr string) error
	EnableDebug(debug bool)

	PersonFindOrCreate(firstName, lastName, email string) (PersonInterface, error)
	PersonCreate(person PersonInterface) error
	PersonCreateWithEmail(fields map[string]string) (PersonInterface, error)
	PersonUpdate(person PersonInterface) error
	PersonUpdateByID(personID string, fields map[string]interface{}) error
	PersonFindByID(personID string) (PersonInterface, error)
	PersonFindByEmail(email string) (PersonInterface, error)
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
