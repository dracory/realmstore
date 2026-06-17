package realmstore

import "errors"

// EmailQueryInterface defines the interface for email query operations.
type EmailQueryInterface interface {
	Validate() error

	IsIDSet() bool
	GetID() string
	SetID(id string) EmailQueryInterface

	IsOwnerIDSet() bool
	GetOwnerID() string
	SetOwnerID(ownerID string) EmailQueryInterface

	IsStatusSet() bool
	GetStatus() string
	SetStatus(status string) EmailQueryInterface

	IsEmailAddressSet() bool
	GetEmailAddress() string
	SetEmailAddress(emailAddress string) EmailQueryInterface

	IsIsDefaultSet() bool
	GetIsDefault() string
	SetIsDefault(isDefault string) EmailQueryInterface

	IsSearchSet() bool
	GetSearch() string
	SetSearch(search string) EmailQueryInterface

	IsDeletedOnlySet() bool
	GetDeletedOnly() bool
	SetDeletedOnly(deletedOnly bool) EmailQueryInterface

	IsWithSoftDeletedSet() bool
	GetWithSoftDeleted() bool
	SetWithSoftDeleted(withSoftDeleted bool) EmailQueryInterface

	IsLimitSet() bool
	GetLimit() int
	SetLimit(limit int) EmailQueryInterface

	IsOffsetSet() bool
	GetOffset() int
	SetOffset(offset int) EmailQueryInterface

	IsSortBySet() bool
	GetSortBy() string
	SetSortBy(sortBy string) EmailQueryInterface

	IsSortOrderSet() bool
	GetSortOrder() string
	SetSortOrder(sortOrder string) EmailQueryInterface
}

// EmailQuery is a shortcut for NewEmailQuery.
func EmailQuery() EmailQueryInterface {
	return NewEmailQuery()
}

// NewEmailQuery creates a new email query.
func NewEmailQuery() EmailQueryInterface {
	return &emailQuery{}
}

var _ EmailQueryInterface = (*emailQuery)(nil)

type emailQuery struct {
	isIDSet              bool
	id                   string
	isOwnerIDSet         bool
	ownerID              string
	isStatusSet          bool
	status               string
	isEmailAddressSet    bool
	emailAddress         string
	isIsDefaultSet       bool
	isDefault            string
	isSearchSet          bool
	search               string
	isDeletedOnlySet     bool
	deletedOnly          bool
	isWithSoftDeletedSet bool
	withSoftDeleted      bool
	isLimitSet           bool
	limit                int
	isOffsetSet          bool
	offset               int
	isSortBySet          bool
	sortBy               string
	isSortOrderSet       bool
	sortOrder            string
}

func (q *emailQuery) Validate() error {
	if q.IsIDSet() && q.GetID() == "" {
		return errors.New("email query: id cannot be empty")
	}
	if q.IsLimitSet() && q.GetLimit() < 0 {
		return errors.New("email query: limit cannot be negative")
	}
	if q.IsOffsetSet() && q.GetOffset() < 0 {
		return errors.New("email query: offset cannot be negative")
	}
	return nil
}

func (q *emailQuery) IsIDSet() bool { return q.isIDSet }
func (q *emailQuery) GetID() string { return q.id }
func (q *emailQuery) SetID(id string) EmailQueryInterface {
	q.isIDSet = true
	q.id = id
	return q
}

func (q *emailQuery) IsOwnerIDSet() bool { return q.isOwnerIDSet }
func (q *emailQuery) GetOwnerID() string { return q.ownerID }
func (q *emailQuery) SetOwnerID(ownerID string) EmailQueryInterface {
	q.isOwnerIDSet = true
	q.ownerID = ownerID
	return q
}

func (q *emailQuery) IsStatusSet() bool { return q.isStatusSet }
func (q *emailQuery) GetStatus() string { return q.status }
func (q *emailQuery) SetStatus(status string) EmailQueryInterface {
	q.isStatusSet = true
	q.status = status
	return q
}

func (q *emailQuery) IsEmailAddressSet() bool { return q.isEmailAddressSet }
func (q *emailQuery) GetEmailAddress() string { return q.emailAddress }
func (q *emailQuery) SetEmailAddress(emailAddress string) EmailQueryInterface {
	q.isEmailAddressSet = true
	q.emailAddress = emailAddress
	return q
}

func (q *emailQuery) IsIsDefaultSet() bool { return q.isIsDefaultSet }
func (q *emailQuery) GetIsDefault() string { return q.isDefault }
func (q *emailQuery) SetIsDefault(isDefault string) EmailQueryInterface {
	q.isIsDefaultSet = true
	q.isDefault = isDefault
	return q
}

func (q *emailQuery) IsSearchSet() bool { return q.isSearchSet }
func (q *emailQuery) GetSearch() string { return q.search }
func (q *emailQuery) SetSearch(search string) EmailQueryInterface {
	q.isSearchSet = true
	q.search = search
	return q
}

func (q *emailQuery) IsDeletedOnlySet() bool { return q.isDeletedOnlySet }
func (q *emailQuery) GetDeletedOnly() bool   { return q.deletedOnly }
func (q *emailQuery) SetDeletedOnly(deletedOnly bool) EmailQueryInterface {
	q.isDeletedOnlySet = true
	q.deletedOnly = deletedOnly
	return q
}

func (q *emailQuery) IsWithSoftDeletedSet() bool { return q.isWithSoftDeletedSet }
func (q *emailQuery) GetWithSoftDeleted() bool   { return q.withSoftDeleted }
func (q *emailQuery) SetWithSoftDeleted(withSoftDeleted bool) EmailQueryInterface {
	q.isWithSoftDeletedSet = true
	q.withSoftDeleted = withSoftDeleted
	return q
}

func (q *emailQuery) IsLimitSet() bool { return q.isLimitSet }
func (q *emailQuery) GetLimit() int    { return q.limit }
func (q *emailQuery) SetLimit(limit int) EmailQueryInterface {
	q.isLimitSet = true
	q.limit = limit
	return q
}

func (q *emailQuery) IsOffsetSet() bool { return q.isOffsetSet }
func (q *emailQuery) GetOffset() int    { return q.offset }
func (q *emailQuery) SetOffset(offset int) EmailQueryInterface {
	q.isOffsetSet = true
	q.offset = offset
	return q
}

func (q *emailQuery) IsSortBySet() bool { return q.isSortBySet }
func (q *emailQuery) GetSortBy() string { return q.sortBy }
func (q *emailQuery) SetSortBy(sortBy string) EmailQueryInterface {
	q.isSortBySet = true
	q.sortBy = sortBy
	return q
}

func (q *emailQuery) IsSortOrderSet() bool { return q.isSortOrderSet }
func (q *emailQuery) GetSortOrder() string { return q.sortOrder }
func (q *emailQuery) SetSortOrder(sortOrder string) EmailQueryInterface {
	q.isSortOrderSet = true
	q.sortOrder = sortOrder
	return q
}
