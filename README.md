# realmstore

A Go library providing a database abstraction layer for managing Person and Email entities with full CRUD operations.

## Features

- **Multi-database support**: MySQL, PostgreSQL, SQLite, MSSQL
- **Auto-migration**: Automatically create tables for Person and Email entities
- **CRUD Operations**: Full Create, Read, Update, Delete functionality
- **Query builders**: Flexible query options with filtering, pagination, and sorting
- **Debug mode**: SQL query logging for development

## Installation

```bash
go get github.com/dracory/realmstore
```

## Quick Start

```go
package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/dracory/realmstore"
)

func main() {
    db, _ := sql.Open("mysql", "user:password@/dbname")
    
    store, err := realmstore.NewStore(
        realmstore.WithDb(db),
        realmstore.WithPersonTableName("persons"),
        realmstore.WithEmailTableName("emails"),
        realmstore.WithAutoMigrate(true),
        realmstore.WithDebug(true),
    )
    if err != nil {
        panic(err)
    }

    // Create a person with email
    person := &realmstore.Person{
        Status:    realmstore.PersonStatusActive,
        FirstName: "John",
        LastName:  "Doe",
    }
    
    err = store.PersonCreate(person)
    // ...
}
```

## Configuration Options

| Option | Description |
|--------|-------------|
| `WithDb(db)` | Set the database connection |
| `WithPersonTableName(name)` | Set the persons table name |
| `WithEmailTableName(name)` | Set the emails table name |
| `WithAutoMigrate(enabled)` | Enable automatic table creation |
| `WithDebug(enabled)` | Enable SQL query logging |

## Person Operations

- `PersonCreate(person)` - Create a new person
- `PersonFindById(id)` - Find person by ID
- `PersonFindByEmail(email)` - Find person by email
- `PersonFindDefaultEmail(id)` - Get person's default email
- `PersonUpdate(person)` - Update a person
- `PersonUpdateById(id, updates)` - Update person by ID
- `PersonDelete(id)` - Delete a person
- `PersonList(opts)` - List persons with pagination
- `PersonFindOrCreate(firstName, lastName, email)` - Find or create person

## Email Operations

- `EmailCreate(email)` - Create a new email
- `EmailFindByID(id)` - Find email by ID
- `EmailFindByEmailAddress(address)` - Find email by address
- `EmailFind(personId)` - Find email by person ID
- `EmailsGetByPersonID(id)` - Get all emails for a person
- `EmailUpdate(email)` - Update an email
- `EmailUpdateByID(id, updates)` - Update email by ID
- `EmailDeleteByID(id)` - Delete email by ID
- `EmailList(opts)` - List emails with pagination

## Models

### Person

| Field | Type | Description |
|-------|------|-------------|
| ID | string | Unique identifier |
| Status | string | Person status (active/inactive) |
| FirstName | string | First name |
| MiddleNames | string | Middle names |
| LastName | string | Last name |
| BirthDay | *time.Time | Date of birth |
| CreatedAt | time.Time | Creation timestamp |
| UpdatedAt | time.Time | Last update timestamp |
| DeletedAt | *time.Time | Deletion timestamp |

### Email

| Field | Type | Description |
|-------|------|-------------|
| ID | string | Unique identifier |
| OwnerID | string | ID of the person who owns this email |
| Status | string | Email status (active/inactive/pending_activation) |
| EmailAddress | string | Email address |
| IsDefault | string | Whether this is the default email (yes/no) |
| CreatedAt | time.Time | Creation timestamp |
| UpdatedAt | time.Time | Last update timestamp |
| DeletedAt | *time.Time | Deletion timestamp |

## Constants

### Person Status
- `PersonStatusActive` - "active"
- `PersonStatusInactive` - "inactive"

### Email Status
- `EmailStatusActive` - "active"
- `EmailStatusInactive` - "inactive"
- `EmailStatusPendingActivation` - "pending_activation"

### Email Type
- `EmailTypePersonal` - "personal"
- `EmailTypeWork` - "work"

## Dependencies

- [goqu](https://github.com/doug-martin/goqu) - SQL query builder
- [scany](https://github.com/georgysavva/scany) - Database scan utility
- [dracory/uid](https://github.com/dracory/uid) - Unique ID generator

## License

MIT
