package realmstore

import (
	"time"
)

type emailImplementation struct {
	ID           string     `json:"id"  db:"id"`
	OwnerID      string     `json:"owner_id" db:"owner_id"`
	Status       string     `json:"status" db:"status"`
	EmailAddress string     `json:"email_address" db:"email_address"`
	IsDefault    string     `json:"is_default" db:"is_default"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
}

var _ EmailInterface = (*emailImplementation)(nil)

func (e *emailImplementation) GetID() string {
	return e.ID
}

func (e *emailImplementation) SetID(id string) {
	e.ID = id
}

func (e *emailImplementation) GetOwnerID() string {
	return e.OwnerID
}

func (e *emailImplementation) SetOwnerID(ownerID string) {
	e.OwnerID = ownerID
}

func (e *emailImplementation) GetStatus() string {
	return e.Status
}

func (e *emailImplementation) SetStatus(status string) {
	e.Status = status
}

func (e *emailImplementation) GetEmailAddress() string {
	return e.EmailAddress
}

func (e *emailImplementation) SetEmailAddress(emailAddress string) {
	e.EmailAddress = emailAddress
}

func (e *emailImplementation) GetIsDefault() string {
	return e.IsDefault
}

func (e *emailImplementation) SetIsDefault(isDefault string) {
	e.IsDefault = isDefault
}

func (e *emailImplementation) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *emailImplementation) SetCreatedAt(createdAt time.Time) {
	e.CreatedAt = createdAt
}

func (e *emailImplementation) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

func (e *emailImplementation) SetUpdatedAt(updatedAt time.Time) {
	e.UpdatedAt = updatedAt
}

func (e *emailImplementation) GetDeletedAt() *time.Time {
	return e.DeletedAt
}

func (e *emailImplementation) SetDeletedAt(deletedAt *time.Time) {
	e.DeletedAt = deletedAt
}

func NewEmail() EmailInterface {
	return &emailImplementation{}
}
