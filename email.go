package realmstore

import (
	"time"

	"github.com/dracory/neat/database/orm"
	"github.com/dracory/neat/database/soft_delete"
	neatuid "github.com/dracory/neat/support/uid"
	"github.com/dromara/carbon/v2"
)

// EmailInterface defines the public API for an email
type EmailInterface interface {
	GetID() string
	SetID(id string) EmailInterface

	GetOwnerID() string
	SetOwnerID(ownerID string) EmailInterface

	GetStatus() string
	SetStatus(status string) EmailInterface

	GetEmailAddress() string
	SetEmailAddress(emailAddress string) EmailInterface

	GetIsDefault() string
	SetIsDefault(isDefault string) EmailInterface

	GetCreatedAt() string
	GetCreatedAtCarbon() *carbon.Carbon
	SetCreatedAt(createdAt string) EmailInterface

	GetUpdatedAt() string
	GetUpdatedAtCarbon() *carbon.Carbon
	SetUpdatedAt(updatedAt string) EmailInterface

	GetSoftDeletedAt() string
	GetSoftDeletedAtCarbon() *carbon.Carbon
	SetSoftDeletedAt(deletedAt string) EmailInterface

	IsSoftDeleted() bool
}

// emailImplementation is the concrete implementation of EmailInterface
type emailImplementation struct {
	orm.ShortID

	OwnerIDField      string `db:"owner_id"`
	StatusField       string `db:"status"`
	EmailAddressField string `db:"email_address"`
	IsDefaultField    string `db:"is_default"`
	CreatedAtField    orm.CreatedAt
	UpdatedAtField    orm.UpdatedAt
	soft_delete.SoftDeletesMaxDate
}

var _ EmailInterface = (*emailImplementation)(nil)

// NewEmail creates a new email.
func NewEmail() EmailInterface {
	o := &emailImplementation{}
	o.SetID(neatuid.GenerateShortID())
	o.SetStatus(EmailStatusActive)
	o.SetEmailAddress("")
	o.SetIsDefault(No)
	o.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString())
	o.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())
	o.SetSoftDeletedAt(MAX_DATETIME)
	return o
}

// NewEmailFromExistingData creates a new email from a raw column map.
func NewEmailFromExistingData(data map[string]string) EmailInterface {
	o := &emailImplementation{}
	o.SetID(data[COLUMN_ID])
	o.SetOwnerID(data[COLUMN_OWNER_ID])
	o.SetStatus(data[COLUMN_STATUS])
	o.SetEmailAddress(data[COLUMN_EMAIL_ADDRESS])
	o.SetIsDefault(data[COLUMN_IS_DEFAULT])
	if v, ok := data[COLUMN_CREATED_AT]; ok {
		o.SetCreatedAt(v)
	}
	if v, ok := data[COLUMN_UPDATED_AT]; ok {
		o.SetUpdatedAt(v)
	}
	if v, ok := data[COLUMN_SOFT_DELETED_AT]; ok {
		o.SetSoftDeletedAt(v)
	}
	return o
}

// == SETTERS AND GETTERS =====================================================

func (e *emailImplementation) GetID() string {
	return e.ShortID.ID
}

func (e *emailImplementation) SetID(id string) EmailInterface {
	e.ShortID.ID = id
	return e
}

func (e *emailImplementation) GetOwnerID() string {
	return e.OwnerIDField
}

func (e *emailImplementation) SetOwnerID(ownerID string) EmailInterface {
	e.OwnerIDField = ownerID
	return e
}

func (e *emailImplementation) GetStatus() string {
	return e.StatusField
}

func (e *emailImplementation) SetStatus(status string) EmailInterface {
	e.StatusField = status
	return e
}

func (e *emailImplementation) GetEmailAddress() string {
	return e.EmailAddressField
}

func (e *emailImplementation) SetEmailAddress(emailAddress string) EmailInterface {
	e.EmailAddressField = emailAddress
	return e
}

func (e *emailImplementation) GetIsDefault() string {
	return e.IsDefaultField
}

func (e *emailImplementation) SetIsDefault(isDefault string) EmailInterface {
	e.IsDefaultField = isDefault
	return e
}

func (e *emailImplementation) GetCreatedAt() string {
	if e.CreatedAtField.CreatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(e.CreatedAtField.CreatedAt).ToDateTimeString()
}

func (e *emailImplementation) GetCreatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(e.CreatedAtField.CreatedAt)
}

func (e *emailImplementation) SetCreatedAt(createdAt string) EmailInterface {
	if createdAt == "" {
		return e
	}
	e.CreatedAtField.CreatedAt = carbon.Parse(createdAt, carbon.UTC).StdTime()
	return e
}

func (e *emailImplementation) GetUpdatedAt() string {
	if e.UpdatedAtField.UpdatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(e.UpdatedAtField.UpdatedAt).ToDateTimeString()
}

func (e *emailImplementation) GetUpdatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(e.UpdatedAtField.UpdatedAt)
}

func (e *emailImplementation) SetUpdatedAt(updatedAt string) EmailInterface {
	if updatedAt == "" {
		return e
	}
	e.UpdatedAtField.UpdatedAt = carbon.Parse(updatedAt, carbon.UTC).StdTime()
	return e
}

func (e *emailImplementation) GetSoftDeletedAt() string {
	if e.SoftDeletedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(e.SoftDeletedAt).ToDateTimeString()
}

func (e *emailImplementation) GetSoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(e.SoftDeletedAt)
}

func (e *emailImplementation) SetSoftDeletedAt(deletedAt string) EmailInterface {
	if deletedAt == "" {
		e.SoftDeletedAt = time.Time{}
		return e
	}
	e.SoftDeletedAt = carbon.Parse(deletedAt, carbon.UTC).StdTime()
	return e
}

func (e *emailImplementation) IsSoftDeleted() bool {
	return e.SoftDeletedAt.Before(carbon.Now(carbon.UTC).StdTime())
}
