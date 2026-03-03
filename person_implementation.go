package realmstore

import (
	"time"
)

type personImplementation struct {
	ID          string     `json:"id" db:"id"`
	Status      string     `json:"status" db:"status"`
	FirstName   string     `json:"first_name" db:"first_name"`
	MiddleNames string     `json:"middle_names" db:"middle_names"`
	LastName    string     `json:"last_name" db:"last_name"`
	BirthDay    *time.Time `json:"birth_day" db:"birth_day"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

var _ PersonInterface = (*personImplementation)(nil)

func (p *personImplementation) GetID() string {
	return p.ID
}

func (p *personImplementation) SetID(id string) {
	p.ID = id
}

func (p *personImplementation) GetStatus() string {
	return p.Status
}

func (p *personImplementation) SetStatus(status string) {
	p.Status = status
}

func (p *personImplementation) GetFirstName() string {
	return p.FirstName
}

func (p *personImplementation) SetFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *personImplementation) GetMiddleNames() string {
	return p.MiddleNames
}

func (p *personImplementation) SetMiddleNames(middleNames string) {
	p.MiddleNames = middleNames
}

func (p *personImplementation) GetLastName() string {
	return p.LastName
}

func (p *personImplementation) SetLastName(lastName string) {
	p.LastName = lastName
}

func (p *personImplementation) GetBirthDay() *time.Time {
	return p.BirthDay
}

func (p *personImplementation) SetBirthDay(birthDay *time.Time) {
	p.BirthDay = birthDay
}

func (p *personImplementation) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *personImplementation) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt
}

func (p *personImplementation) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

func (p *personImplementation) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt
}

func (p *personImplementation) GetDeletedAt() *time.Time {
	return p.DeletedAt
}

func (p *personImplementation) SetDeletedAt(deletedAt *time.Time) {
	p.DeletedAt = deletedAt
}

func NewPerson() PersonInterface {
	return &personImplementation{}
}
