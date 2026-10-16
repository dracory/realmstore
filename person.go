package realmstore

import (
	"time"
	"uuid"

	"github.com/dracory/neat/database/orm"
	"github.com/dracory/neat/database/soft_delete"
	"github.com/dromara/carbon/v2"
)

// PersonInterface defines the public API for a person
type PersonInterface interface {
	GetID() string
	SetID(id string) PersonInterface

	GetStatus() string
	SetStatus(status string) PersonInterface

	GetFirstName() string
	SetFirstName(firstName string) PersonInterface

	GetMiddleNames() string
	SetMiddleNames(middleNames string) PersonInterface

	GetLastName() string
	SetLastName(lastName string) PersonInterface

	GetBirthDay() *time.Time
	SetBirthDay(birthDay *time.Time) PersonInterface

	GetCreatedAt() string
	GetCreatedAtCarbon() *carbon.Carbon
	SetCreatedAt(createdAt string) PersonInterface

	GetUpdatedAt() string
	GetUpdatedAtCarbon() *carbon.Carbon
	SetUpdatedAt(updatedAt string) PersonInterface

	GetSoftDeletedAt() string
	GetSoftDeletedAtCarbon() *carbon.Carbon
	SetSoftDeletedAt(deletedAt string) PersonInterface

	IsSoftDeleted() bool
}

// personImplementation is the concrete implementation of PersonInterface
type personImplementation struct {
	orm.StringID

	StatusField      string     `db:"status"`
	FirstNameField   string     `db:"first_name"`
	MiddleNamesField string     `db:"middle_names"`
	LastNameField    string     `db:"last_name"`
	BirthDayField    *time.Time `db:"birth_day"`
	CreatedAtField   orm.CreatedAt
	UpdatedAtField   orm.UpdatedAt
	soft_delete.SoftDeletesMaxDate
}

var _ PersonInterface = (*personImplementation)(nil)

// NewPerson creates a new person.
func NewPerson() PersonInterface {
	o := &personImplementation{}
	o.SetID(uuid.NewV7().String())
	o.SetStatus(PersonStatusActive)
	o.SetFirstName("")
	o.SetMiddleNames("")
	o.SetLastName("")
	o.SetBirthDay(nil)
	o.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString())
	o.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())
	o.SetSoftDeletedAt(MAX_DATETIME)
	return o
}

// NewPersonFromExistingData creates a new person from a raw column map.
func NewPersonFromExistingData(data map[string]string) PersonInterface {
	o := &personImplementation{}
	o.SetID(data[COLUMN_ID])
	o.SetStatus(data[COLUMN_STATUS])
	o.SetFirstName(data[COLUMN_FIRST_NAME])
	o.SetMiddleNames(data[COLUMN_MIDDLE_NAMES])
	o.SetLastName(data[COLUMN_LAST_NAME])
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

func (p *personImplementation) GetID() string {
	return p.StringID.ID
}

func (p *personImplementation) SetID(id string) PersonInterface {
	p.StringID.ID = id
	return p
}

func (p *personImplementation) GetStatus() string {
	return p.StatusField
}

func (p *personImplementation) SetStatus(status string) PersonInterface {
	p.StatusField = status
	return p
}

func (p *personImplementation) GetFirstName() string {
	return p.FirstNameField
}

func (p *personImplementation) SetFirstName(firstName string) PersonInterface {
	p.FirstNameField = firstName
	return p
}

func (p *personImplementation) GetMiddleNames() string {
	return p.MiddleNamesField
}

func (p *personImplementation) SetMiddleNames(middleNames string) PersonInterface {
	p.MiddleNamesField = middleNames
	return p
}

func (p *personImplementation) GetLastName() string {
	return p.LastNameField
}

func (p *personImplementation) SetLastName(lastName string) PersonInterface {
	p.LastNameField = lastName
	return p
}

func (p *personImplementation) GetBirthDay() *time.Time {
	return p.BirthDayField
}

func (p *personImplementation) SetBirthDay(birthDay *time.Time) PersonInterface {
	p.BirthDayField = birthDay
	return p
}

func (p *personImplementation) GetCreatedAt() string {
	if p.CreatedAtField.CreatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(p.CreatedAtField.CreatedAt).ToDateTimeString()
}

func (p *personImplementation) GetCreatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(p.CreatedAtField.CreatedAt)
}

func (p *personImplementation) SetCreatedAt(createdAt string) PersonInterface {
	if createdAt == "" {
		return p
	}
	p.CreatedAtField.CreatedAt = carbon.Parse(createdAt, carbon.UTC).StdTime()
	return p
}

func (p *personImplementation) GetUpdatedAt() string {
	if p.UpdatedAtField.UpdatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(p.UpdatedAtField.UpdatedAt).ToDateTimeString()
}

func (p *personImplementation) GetUpdatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(p.UpdatedAtField.UpdatedAt)
}

func (p *personImplementation) SetUpdatedAt(updatedAt string) PersonInterface {
	if updatedAt == "" {
		return p
	}
	p.UpdatedAtField.UpdatedAt = carbon.Parse(updatedAt, carbon.UTC).StdTime()
	return p
}

func (p *personImplementation) GetSoftDeletedAt() string {
	if p.SoftDeletedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(p.SoftDeletedAt).ToDateTimeString()
}

func (p *personImplementation) GetSoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(p.SoftDeletedAt)
}

func (p *personImplementation) SetSoftDeletedAt(deletedAt string) PersonInterface {
	if deletedAt == "" {
		p.SoftDeletedAt = time.Time{}
		return p
	}
	p.SoftDeletedAt = carbon.Parse(deletedAt, carbon.UTC).StdTime()
	return p
}

func (p *personImplementation) IsSoftDeleted() bool {
	return p.SoftDeletedAt.Before(carbon.Now(carbon.UTC).StdTime())
}

// PersonWithEmailInterface defines the public API for a person with their email
type PersonWithEmailInterface interface {
	PersonInterface
	GetEmailAddress() string
	SetEmailAddress(emailAddress string) PersonWithEmailInterface
}

// personWithEmail is a person with their email address attached.
type personWithEmail struct {
	personImplementation
	EmailAddressField string `db:"email_address"`
}

var _ PersonWithEmailInterface = (*personWithEmail)(nil)

func (p *personWithEmail) GetEmailAddress() string {
	return p.EmailAddressField
}

func (p *personWithEmail) SetEmailAddress(emailAddress string) PersonWithEmailInterface {
	p.EmailAddressField = emailAddress
	return p
}
