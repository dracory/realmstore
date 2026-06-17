package realmstore

import "errors"

// PersonQueryInterface defines the interface for person query operations.
type PersonQueryInterface interface {
	Validate() error

	IsIDSet() bool
	GetID() string
	SetID(id string) PersonQueryInterface

	IsStatusSet() bool
	GetStatus() string
	SetStatus(status string) PersonQueryInterface

	IsSearchSet() bool
	GetSearch() string
	SetSearch(search string) PersonQueryInterface

	IsWithEmailSet() bool
	GetWithEmail() bool
	SetWithEmail(withEmail bool) PersonQueryInterface

	IsDeletedOnlySet() bool
	GetDeletedOnly() bool
	SetDeletedOnly(deletedOnly bool) PersonQueryInterface

	IsWithSoftDeletedSet() bool
	GetWithSoftDeleted() bool
	SetWithSoftDeleted(withSoftDeleted bool) PersonQueryInterface

	IsLimitSet() bool
	GetLimit() int
	SetLimit(limit int) PersonQueryInterface

	IsOffsetSet() bool
	GetOffset() int
	SetOffset(offset int) PersonQueryInterface

	IsSortBySet() bool
	GetSortBy() string
	SetSortBy(sortBy string) PersonQueryInterface

	IsSortOrderSet() bool
	GetSortOrder() string
	SetSortOrder(sortOrder string) PersonQueryInterface
}

// PersonQuery is a shortcut for NewPersonQuery.
func PersonQuery() PersonQueryInterface {
	return NewPersonQuery()
}

// NewPersonQuery creates a new person query.
func NewPersonQuery() PersonQueryInterface {
	return &personQuery{}
}

var _ PersonQueryInterface = (*personQuery)(nil)

type personQuery struct {
	isIDSet              bool
	id                   string
	isStatusSet          bool
	status               string
	isSearchSet          bool
	search               string
	isWithEmailSet       bool
	withEmail            bool
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

func (q *personQuery) Validate() error {
	if q.IsIDSet() && q.GetID() == "" {
		return errors.New("person query: id cannot be empty")
	}
	if q.IsLimitSet() && q.GetLimit() < 0 {
		return errors.New("person query: limit cannot be negative")
	}
	if q.IsOffsetSet() && q.GetOffset() < 0 {
		return errors.New("person query: offset cannot be negative")
	}
	return nil
}

func (q *personQuery) IsIDSet() bool { return q.isIDSet }
func (q *personQuery) GetID() string { return q.id }
func (q *personQuery) SetID(id string) PersonQueryInterface {
	q.isIDSet = true
	q.id = id
	return q
}

func (q *personQuery) IsStatusSet() bool { return q.isStatusSet }
func (q *personQuery) GetStatus() string { return q.status }
func (q *personQuery) SetStatus(status string) PersonQueryInterface {
	q.isStatusSet = true
	q.status = status
	return q
}

func (q *personQuery) IsSearchSet() bool { return q.isSearchSet }
func (q *personQuery) GetSearch() string { return q.search }
func (q *personQuery) SetSearch(search string) PersonQueryInterface {
	q.isSearchSet = true
	q.search = search
	return q
}

func (q *personQuery) IsWithEmailSet() bool { return q.isWithEmailSet }
func (q *personQuery) GetWithEmail() bool   { return q.withEmail }
func (q *personQuery) SetWithEmail(withEmail bool) PersonQueryInterface {
	q.isWithEmailSet = true
	q.withEmail = withEmail
	return q
}

func (q *personQuery) IsDeletedOnlySet() bool { return q.isDeletedOnlySet }
func (q *personQuery) GetDeletedOnly() bool   { return q.deletedOnly }
func (q *personQuery) SetDeletedOnly(deletedOnly bool) PersonQueryInterface {
	q.isDeletedOnlySet = true
	q.deletedOnly = deletedOnly
	return q
}

func (q *personQuery) IsWithSoftDeletedSet() bool { return q.isWithSoftDeletedSet }
func (q *personQuery) GetWithSoftDeleted() bool   { return q.withSoftDeleted }
func (q *personQuery) SetWithSoftDeleted(withSoftDeleted bool) PersonQueryInterface {
	q.isWithSoftDeletedSet = true
	q.withSoftDeleted = withSoftDeleted
	return q
}

func (q *personQuery) IsLimitSet() bool { return q.isLimitSet }
func (q *personQuery) GetLimit() int    { return q.limit }
func (q *personQuery) SetLimit(limit int) PersonQueryInterface {
	q.isLimitSet = true
	q.limit = limit
	return q
}

func (q *personQuery) IsOffsetSet() bool { return q.isOffsetSet }
func (q *personQuery) GetOffset() int    { return q.offset }
func (q *personQuery) SetOffset(offset int) PersonQueryInterface {
	q.isOffsetSet = true
	q.offset = offset
	return q
}

func (q *personQuery) IsSortBySet() bool { return q.isSortBySet }
func (q *personQuery) GetSortBy() string { return q.sortBy }
func (q *personQuery) SetSortBy(sortBy string) PersonQueryInterface {
	q.isSortBySet = true
	q.sortBy = sortBy
	return q
}

func (q *personQuery) IsSortOrderSet() bool { return q.isSortOrderSet }
func (q *personQuery) GetSortOrder() string { return q.sortOrder }
func (q *personQuery) SetSortOrder(sortOrder string) PersonQueryInterface {
	q.isSortOrderSet = true
	q.sortOrder = sortOrder
	return q
}
