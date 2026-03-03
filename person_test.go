package realmstore

import (
	"testing"
	"time"
)

func TestPersonCreate(t *testing.T) {
	store := newTestStore(t)

	person := &Person{
		Status:    PersonStatusActive,
		FirstName: "John",
		LastName:  "Doe",
	}

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	if person.ID == "" {
		t.Error("expected person ID to be set")
	}

	if person.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set")
	}

	if person.UpdatedAt.IsZero() {
		t.Error("expected UpdatedAt to be set")
	}
}

func TestPersonFindByID(t *testing.T) {
	store := newTestStore(t)

	person := &Person{
		Status:    PersonStatusActive,
		FirstName: "Jane",
		LastName:  "Smith",
	}

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	found, err := store.PersonFindByID(person.ID)
	if err != nil {
		t.Fatalf("failed to find person: %v", err)
	}

	if found == nil {
		t.Fatal("expected to find person")
	}

	if found.FirstName != "Jane" {
		t.Errorf("expected FirstName to be 'Jane', got '%s'", found.FirstName)
	}

	if found.LastName != "Smith" {
		t.Errorf("expected LastName to be 'Smith', got '%s'", found.LastName)
	}
}

func TestPersonFindByIDNotFound(t *testing.T) {
	store := newTestStore(t)

	found, err := store.PersonFindByID("nonexistent-id")
	if err != nil {
		t.Fatalf("failed to find person: %v", err)
	}

	if found != nil {
		t.Error("expected nil person for nonexistent ID")
	}
}

func TestPersonUpdate(t *testing.T) {
	store := newTestStore(t)

	person := &Person{
		Status:    PersonStatusActive,
		FirstName: "Bob",
		LastName:  "Wilson",
	}

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	originalUpdatedAt := person.UpdatedAt
	time.Sleep(10 * time.Millisecond)

	person.FirstName = "Robert"
	err = store.PersonUpdate(person)
	if err != nil {
		t.Fatalf("failed to update person: %v", err)
	}

	if !person.UpdatedAt.After(originalUpdatedAt) {
		t.Error("expected UpdatedAt to be changed")
	}

	found, err := store.PersonFindByID(person.ID)
	if err != nil {
		t.Fatalf("failed to find person: %v", err)
	}

	if found.FirstName != "Robert" {
		t.Errorf("expected FirstName to be 'Robert', got '%s'", found.FirstName)
	}
}

func TestPersonDelete(t *testing.T) {
	store := newTestStore(t)

	person := &Person{
		Status:    PersonStatusActive,
		FirstName: "ToDelete",
		LastName:  "User",
	}

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	err = store.PersonDelete(person.ID)
	if err != nil {
		t.Fatalf("failed to delete person: %v", err)
	}

	found, err := store.PersonFindByID(person.ID)
	if err != nil {
		t.Fatalf("failed to find person: %v", err)
	}

	if found != nil {
		t.Error("expected nil person after delete")
	}
}

func TestPersonFindOrCreate(t *testing.T) {
	store := newTestStore(t)

	person, err := store.PersonFindOrCreate("John", "Doe", "john@example.com")
	if err != nil {
		t.Fatalf("failed to find or create person: %v", err)
	}

	if person.ID == "" {
		t.Error("expected person ID to be set")
	}

	if person.FirstName != "John" {
		t.Errorf("expected FirstName to be 'John', got '%s'", person.FirstName)
	}

	if person.LastName != "Doe" {
		t.Errorf("expected LastName to be 'Doe', got '%s'", person.LastName)
	}

	person2, err := store.PersonFindOrCreate("John", "Doe", "john@example.com")
	if err != nil {
		t.Fatalf("failed to find or create person: %v", err)
	}

	if person2.ID != person.ID {
		t.Error("expected same person to be returned for existing email")
	}
}
