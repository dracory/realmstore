package realmstore

import (
	"time"
)

type PersonInterface interface {
	GetID() string
	SetID(id string)
	GetStatus() string
	SetStatus(status string)
	GetFirstName() string
	SetFirstName(firstName string)
	GetMiddleNames() string
	SetMiddleNames(middleNames string)
	GetLastName() string
	SetLastName(lastName string)
	GetBirthDay() *time.Time
	SetBirthDay(birthDay *time.Time)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	GetDeletedAt() *time.Time
	SetDeletedAt(deletedAt *time.Time)
}
