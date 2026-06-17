# Realm Store

A Go library providing a database abstraction layer for managing **Person** and **Email** entities with full CRUD operations, powered by the [neat](https://github.com/dracory/neat) ORM.

## Features

- **Neat ORM integration**: Uses `dracory/neat` for schema migrations, query building, and soft deletes
- **Auto-migration**: Automatically create tables for Person and Email entities
- **CRUD Operations**: Full Create, Read, Update, Delete functionality
- **Query builders**: Flexible query interfaces with filtering, pagination, sorting, and soft-delete handling
- **Debug mode**: SQL query logging for development
- **Soft deletes**: Records use a far-future sentinel (`MAX_DATETIME`) instead of hard deletion

## Installation

```bash
go get github.com/dracory/realmstore
```

## Quick Start

```go
package main

import (
    "context"
    "database/sql"
    "log"

    _ "modernc.org/sqlite"
    "github.com/dracory/realmstore"
)

func main() {
    db, err := sql.Open("sqlite", "file:test.db?cache=shared")
    if err != nil {
        log.Fatal(err)
    }

    store, err := realmstore.NewStore(realmstore.NewStoreOptions{
        DB:                 db,
        PersonTableName:    "persons",
        EmailTableName:     "emails",
        AutomigrateEnabled: true,
        DebugEnabled:       false,
    })
    if err != nil {
        log.Fatal(err)
    }

    // Create a person
    person := realmstore.NewPerson()
    person.SetStatus(realmstore.PersonStatusActive)
    person.SetFirstName("John")
    person.SetLastName("Doe")

    err = store.PersonCreate(context.Background(), person)
    if err != nil {
        log.Fatal(err)
    }

    // Create an email for the person
    email := realmstore.NewEmail()
    email.SetOwnerID(person.GetID())
    email.SetEmailAddress("john@example.com")
    email.SetIsDefault(realmstore.Yes)

    err = store.EmailCreate(context.Background(), email)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Configuration Options

`NewStoreOptions` fields:

| Field | Type | Description |
|-------|------|-------------|
| `DB` | `*sql.DB` | Database connection (required) |
| `PersonTableName` | `string` | Persons table name (required) |
| `EmailTableName` | `string` | Emails table name (required) |
| `AutomigrateEnabled` | `bool` | Auto-create tables on startup |
| `DebugEnabled` | `bool` | Enable SQL query logging |

## Person Operations

All methods accept `context.Context` as the first argument.

- `PersonCreate(ctx, person)` - Create a new person
- `PersonCreateWithEmail(ctx, fields)` - Create a person and their email in one call
- `PersonFindByID(ctx, id)` - Find person by ID
- `PersonFindByEmail(ctx, email)` - Find person by email address
- `PersonFindDefaultEmail(ctx, personID)` - Get a person's default email
- `PersonFindOrCreate(ctx, firstName, lastName, email)` - Find or create a person
- `PersonUpdate(ctx, person)` - Update a person
- `PersonUpdateByID(ctx, id, fields)` - Update person fields by ID
- `PersonDelete(ctx, id)` - Soft-delete a person
- `PersonList(ctx, query)` - List persons with filtering/pagination
- `PersonCount(ctx, query)` - Count matching persons
- `PersonMaxPages(ctx, query)` - Calculate max page count

### Person Query

```go
query := realmstore.PersonQuery().
    SetStatus(realmstore.PersonStatusActive).
    SetSearch("John").
    SetSortBy("first_name").
    SetSortOrder("asc").
    SetLimit(10).
    SetOffset(0)

list, err := store.PersonList(ctx, query)
count, err := store.PersonCount(ctx, query)
```

Query options:

| Method | Description |
|--------|-------------|
| `SetID(id)` | Filter by ID |
| `SetStatus(status)` | Filter by status |
| `SetSearch(search)` | Search first/last name and email |
| `SetWithEmail(true)` | Include email in results |
| `SetDeletedOnly(true)` | Only soft-deleted records |
| `SetWithSoftDeleted(true)` | Include soft-deleted records |
| `SetLimit(n)` | Page size |
| `SetOffset(n)` | Skip N records |
| `SetSortBy(column)` | Sort column |
| `SetSortOrder("asc"/"desc")` | Sort direction |

## Email Operations

All methods accept `context.Context` as the first argument.

- `EmailCreate(ctx, email)` - Create a new email
- `EmailFind(ctx, query)` - Find a single email by query
- `EmailFindByID(ctx, id)` - Find email by ID
- `EmailFindByEmailAddress(ctx, address)` - Find email by address
- `EmailFindByPersonID(ctx, personID)` - Find email by owner ID
- `EmailUpdate(ctx, email)` - Update an email
- `EmailUpdateByID(ctx, id, fields)` - Update email fields by ID
- `EmailDeleteByID(ctx, id)` - Soft-delete an email
- `EmailList(ctx, query)` - List emails with filtering/pagination
- `EmailCount(ctx, query)` - Count matching emails

### Email Query

```go
query := realmstore.EmailQuery().
    SetOwnerID(personID).
    SetStatus(realmstore.EmailStatusActive).
    SetIsDefault(realmstore.Yes).
    SetLimit(10)

list, err := store.EmailList(ctx, query)
```

Query options:

| Method | Description |
|--------|-------------|
| `SetID(id)` | Filter by ID |
| `SetOwnerID(id)` | Filter by person ID |
| `SetStatus(status)` | Filter by status |
| `SetEmailAddress(address)` | Filter by exact email address |
| `SetIsDefault("yes"/"no")` | Filter by default flag |
| `SetSearch(search)` | Search email address (LIKE) |
| `SetDeletedOnly(true)` | Only soft-deleted records |
| `SetWithSoftDeleted(true)` | Include soft-deleted records |
| `SetLimit(n)` | Page size |
| `SetOffset(n)` | Skip N records |
| `SetSortBy(column)` | Sort column |
| `SetSortOrder("asc"/"desc")` | Sort direction |

## Models

### PersonInterface

Created via `realmstore.NewPerson()`:

```go
person := realmstore.NewPerson()
person.SetStatus(realmstore.PersonStatusActive)
person.SetFirstName("John")
person.SetLastName("Doe")
person.SetMiddleNames("Robert")
person.SetBirthDay(&birthDate)
```

| Method | Returns | Description |
|--------|---------|-------------|
| `GetID()` / `SetID(id)` | `string` | Unique identifier |
| `GetStatus()` / `SetStatus(status)` | `string` | Person status |
| `GetFirstName()` / `SetFirstName(name)` | `string` | First name |
| `GetMiddleNames()` / `SetMiddleNames(names)` | `string` | Middle names |
| `GetLastName()` / `SetLastName(name)` | `string` | Last name |
| `GetBirthDay()` / `SetBirthDay(t)` | `*time.Time` | Date of birth |
| `GetCreatedAt()` / `SetCreatedAt(s)` | `string` | Creation timestamp (carbon) |
| `GetCreatedAtCarbon()` | `*carbon.Carbon` | Creation as Carbon |
| `GetUpdatedAt()` / `SetUpdatedAt(s)` | `string` | Update timestamp (carbon) |
| `GetUpdatedAtCarbon()` | `*carbon.Carbon` | Updated as Carbon |
| `GetSoftDeletedAt()` / `SetSoftDeletedAt(s)` | `string` | Soft-delete timestamp |
| `GetSoftDeletedAtCarbon()` | `*carbon.Carbon` | Soft-delete as Carbon |
| `IsSoftDeleted()` | `bool` | Whether record is soft-deleted |

### EmailInterface

Created via `realmstore.NewEmail()`:

```go
email := realmstore.NewEmail()
email.SetOwnerID(person.GetID())
email.SetEmailAddress("john@example.com")
email.SetIsDefault(realmstore.Yes)
```

| Method | Returns | Description |
|--------|---------|-------------|
| `GetID()` / `SetID(id)` | `string` | Unique identifier |
| `GetOwnerID()` / `SetOwnerID(id)` | `string` | Person ID who owns this email |
| `GetStatus()` / `SetStatus(status)` | `string` | Email status |
| `GetEmailAddress()` / `SetEmailAddress(addr)` | `string` | Email address |
| `GetIsDefault()` / `SetIsDefault(v)` | `string` | `Yes` or `No` |
| `GetCreatedAt()` / `SetCreatedAt(s)` | `string` | Creation timestamp |
| `GetCreatedAtCarbon()` | `*carbon.Carbon` | Creation as Carbon |
| `GetUpdatedAt()` / `SetUpdatedAt(s)` | `string` | Update timestamp |
| `GetUpdatedAtCarbon()` | `*carbon.Carbon` | Updated as Carbon |
| `GetSoftDeletedAt()` / `SetSoftDeletedAt(s)` | `string` | Soft-delete timestamp |
| `GetSoftDeletedAtCarbon()` | `*carbon.Carbon` | Soft-delete as Carbon |
| `IsSoftDeleted()` | `bool` | Whether record is soft-deleted |

## Constants

### Person Status
- `PersonStatusActive` - `"active"`
- `PersonStatusInactive` - `"inactive"`

### Email Status
- `EmailStatusActive` - `"active"`
- `EmailStatusInactive` - `"inactive"`
- `EmailStatusPendingActivation` - `"pending_activation"`

### Boolean Flags
- `Yes` - `"yes"`
- `No` - `"no"`

## Soft Deletes

Both Person and Email use soft deletes. Records are not physically removed from the database. Instead:

- `PersonDelete` / `EmailDeleteByID` set `soft_deleted_at` to the current time
- By default, `List` and `Find` queries exclude soft-deleted records
- Use `SetWithSoftDeleted(true)` on queries to include them
- Use `SetDeletedOnly(true)` to return only soft-deleted records
- New records are initialized with `soft_deleted_at = MAX_DATETIME` (far future)

## Dependencies

- [neat](https://github.com/dracory/neat) - Database ORM and schema builder
- [carbon](https://github.com/dromara/carbon) - DateTime library for Go

## License

MIT
