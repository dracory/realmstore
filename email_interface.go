package realmstore

import (
	"time"
)

type EmailInterface interface {
	GetID() string
	SetID(id string)
	GetOwnerID() string
	SetOwnerID(ownerID string)
	GetStatus() string
	SetStatus(status string)
	GetEmailAddress() string
	SetEmailAddress(emailAddress string)
	GetIsDefault() string
	SetIsDefault(isDefault string)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	GetDeletedAt() *time.Time
	SetDeletedAt(deletedAt *time.Time)
}
