package realmstore

// Person constants
const (
	// PersonStatusActive status active
	PersonStatusActive = "active"
	// PersonStatusInactive status inactive
	PersonStatusInactive = "inactive"
	// PersonRowsPerPage return how many person will have on the page
	PersonRowsPerPage uint = 20
)

// Email constants
const (
	// EmailStatusActive active status
	EmailStatusActive = "active"
	// EmailStatusInactive inactive status
	EmailStatusInactive = "inactive"
	// EmailStatusPendingActivation pending_activation status
	EmailStatusPendingActivation = "pending_activation"

	// EmailTypePersonal represents a personal type
	EmailTypePersonal = "personal"
	// EmailTypeWork represents a work type
	EmailTypeWork = "work"

	// No represents a boolean no
	No = "no"

	// Yes represents a boolean yes
	Yes = "yes"
)

// Column names for the person table
const (
	COLUMN_BIRTH_DAY     = "birth_day"
	COLUMN_CREATED_AT    = "created_at"
	COLUMN_EMAIL_ADDRESS = "email_address"
	COLUMN_FIRST_NAME    = "first_name"
	COLUMN_ID            = "id"
	COLUMN_IS_DEFAULT    = "is_default"
	COLUMN_LAST_NAME     = "last_name"
	COLUMN_MIDDLE_NAMES  = "middle_names"
	COLUMN_OWNER_ID      = "owner_id"
	COLUMN_SOFT_DELETED_AT = "soft_deleted_at"
	COLUMN_STATUS        = "status"
	COLUMN_UPDATED_AT    = "updated_at"
)

// MAX_DATETIME is a far-future datetime used as the default soft-delete sentinel.
const MAX_DATETIME = "9999-12-31 23:59:59"
