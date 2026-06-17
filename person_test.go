package realmstore

import (
	"context"
	"testing"
)

func TestPersonCreate(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("John")
	person.SetLastName("Doe")

	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	if person.GetID() == "" {
		t.Error("expected person ID to be set")
	}

	if person.GetCreatedAt() == "" {
		t.Error("expected CreatedAt to be set")
	}

	if person.GetUpdatedAt() == "" {
		t.Error("expected UpdatedAt to be set")
	}
}

func TestPersonFindByID(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Jane")
	person.SetLastName("Smith")

	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	found, err := store.PersonFindByID(context.Background(), person.GetID())
	if err != nil {
		t.Fatalf("failed to find person: %v", err)
	}

	if found == nil {
		t.Fatal("expected to find person")
	}

	if found.GetFirstName() != "Jane" {
		t.Errorf("expected FirstName to be 'Jane', got '%s'", found.GetFirstName())
	}

	if found.GetLastName() != "Smith" {
		t.Errorf("expected LastName to be 'Smith', got '%s'", found.GetLastName())
	}
}

func TestPersonFindByIDNotFound(t *testing.T) {
	store := newTestStore(t)

	found, err := store.PersonFindByID(context.Background(), "nonexistent-id")
	if err != nil {
		t.Fatalf("failed to find person: %v", err)
	}

	if found != nil {
		t.Error("expected nil person for nonexistent ID")
	}
}

func TestPersonUpdate(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Bob")
	person.SetLastName("Wilson")

	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	person.SetFirstName("Robert")
	err = store.PersonUpdate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to update person: %v", err)
	}

	found, err := store.PersonFindByID(context.Background(), person.GetID())
	if err != nil {
		t.Fatalf("failed to find person: %v", err)
	}

	if found.GetFirstName() != "Robert" {
		t.Errorf("expected FirstName to be 'Robert', got '%s'", found.GetFirstName())
	}
}

func TestPersonDelete(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("ToDelete")
	person.SetLastName("User")

	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	err = store.PersonDelete(context.Background(), person.GetID())
	if err != nil {
		t.Fatalf("failed to delete person: %v", err)
	}

	found, err := store.PersonFindByID(context.Background(), person.GetID())
	if err != nil {
		t.Fatalf("failed to find person: %v", err)
	}

	if found != nil {
		t.Error("expected nil person after delete")
	}
}

func TestPersonFindOrCreate(t *testing.T) {
	store := newTestStore(t)

	person, err := store.PersonFindOrCreate(context.Background(), "John", "Doe", "john@example.com")
	if err != nil {
		t.Fatalf("failed to find or create person: %v", err)
	}

	if person.GetID() == "" {
		t.Error("expected person ID to be set")
	}

	if person.GetFirstName() != "John" {
		t.Errorf("expected FirstName to be 'John', got '%s'", person.GetFirstName())
	}

	if person.GetLastName() != "Doe" {
		t.Errorf("expected LastName to be 'Doe', got '%s'", person.GetLastName())
	}

	person2, err := store.PersonFindOrCreate(context.Background(), "John", "Doe", "john@example.com")
	if err != nil {
		t.Fatalf("failed to find or create person: %v", err)
	}

	if person2.GetID() != person.GetID() {
		t.Error("expected same person to be returned for existing email")
	}
}
