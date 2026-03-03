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

var _ PersonInterface = (*Person)(nil)

func (p *Person) GetID() string {
	return p.ID
}

func (p *Person) SetID(id string) {
	p.ID = id
}

func (p *Person) GetStatus() string {
	return p.Status
}

func (p *Person) SetStatus(status string) {
	p.Status = status
}

func (p *Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) SetFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *Person) GetMiddleNames() string {
	return p.MiddleNames
}

func (p *Person) SetMiddleNames(middleNames string) {
	p.MiddleNames = middleNames
}

func (p *Person) GetLastName() string {
	return p.LastName
}

func (p *Person) SetLastName(lastName string) {
	p.LastName = lastName
}

func (p *Person) GetBirthDay() *time.Time {
	return p.BirthDay
}

func (p *Person) SetBirthDay(birthDay *time.Time) {
	p.BirthDay = birthDay
}

func (p *Person) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *Person) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt
}

func (p *Person) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

func (p *Person) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt
}

func (p *Person) GetDeletedAt() *time.Time {
	return p.DeletedAt
}

func (p *Person) SetDeletedAt(deletedAt *time.Time) {
	p.DeletedAt = deletedAt
}
