package realmstore

type EmailQueryOptions struct {
	Search      string
	CreatedAt   string
	OwnerID     string
	Status      string
	IsDefault   string // one of: empty, yes, no
	SortBy      string
	SortOrder   string // asc / dec
	Limit       uint
	Offset      uint
	CountOnly   bool // returns only the count if enabled
	DeletedOnly bool // returns only deleted Persons
}
