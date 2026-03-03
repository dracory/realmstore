package realmstore

type PersonQueryOptions struct {
	Search  string
	Created_at  string
	SortBy      string
	SortOrder   string // asc / dec
	Limit       uint
	Offset      uint
	WithEmail   bool // returns Person with his email
	CountOnly   bool // returns only the count if enabled
	DeletedOnly bool // returns only deleted Persons
}
