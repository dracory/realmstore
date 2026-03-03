package realmstore

import (
	"database/sql"
)

type StoreInterface interface {
	AutoMigrate() error
	DriverName(db *sql.DB) string
	DbExec(sqlStr string) error
	EnableDebug(debug bool)

	PersonCreate(person PersonInterface) error
	PersonCreateWithEmail(fields map[string]string) (PersonInterface, error)
	PersonFindOrCreate(firstName, lastName, email string) (PersonInterface, error)
	PersonFindByID(personID string) (PersonInterface, error)
	PersonFindByEmail(email string) (PersonInterface, error)
	PersonFindDefaultEmail(person_id string) (EmailInterface, error)
	PersonDelete(personId string) error
	PersonList(options *PersonQueryOptions) ([]PersonWithEmail, error)
	PersonMaxPages(options *PersonQueryOptions) (uint, error)
	PersonUpdate(person PersonInterface) error
	PersonUpdateByID(personID string, fields map[string]interface{}) error

	EmailCreate(email EmailInterface) (bool, error)
	EmailDeleteByID(emailId string) (bool, error)
	EmailFind(options *EmailQueryOptions) (EmailInterface, error)
	EmailFindByID(id string) (EmailInterface, error)
	EmailFindByEmailAddress(emailAddress string) EmailInterface
	EmailFindByPersonId(personId string) (EmailInterface, error)
	EmailList(options *EmailQueryOptions) ([]EmailInterface, error)
	EmailUpdate(email EmailInterface) bool
	EmailUpdateByID(emailId string, fields map[string]interface{}) bool
}
