package realmstore

import (
	"testing"
)

func TestEmailCreate(t *testing.T) {
	store := newTestStore(t)

	person := &Person{
		Status:    PersonStatusActive,
		FirstName: "Email",
		LastName:  "Owner",
	}

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := &Email{
		OwnerID:      person.ID,
		Status:       EmailStatusActive,
		EmailAddress: "test@example.com",
		IsDefault:    Yes,
	}

	created, err := store.EmailCreate(email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	if !created {
		t.Error("expected email to be created")
	}

	if email.ID == "" {
		t.Error("expected email ID to be set")
	}
}

func TestEmailFindByID(t *testing.T) {
	store := newTestStore(t)

	person := &Person{
		Status:    PersonStatusActive,
		FirstName: "Email",
		LastName:  "Finder",
	}

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := &Email{
		OwnerID:      person.ID,
		Status:       EmailStatusActive,
		EmailAddress: "findme@example.com",
		IsDefault:    Yes,
	}

	_, err = store.EmailCreate(email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	found, err := store.EmailFindByID(email.ID)
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}

	if found == nil {
		t.Fatal("expected to find email")
	}

	if found.EmailAddress != "findme@example.com" {
		t.Errorf("expected EmailAddress to be 'findme@example.com', got '%s'", found.EmailAddress)
	}
}

func TestEmailFindByEmailAddress(t *testing.T) {
	store := newTestStore(t)

	person := &Person{
		Status:    PersonStatusActive,
		FirstName: "Email",
		LastName:  "Address",
	}

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := &Email{
		OwnerID:      person.ID,
		Status:       EmailStatusActive,
		EmailAddress: "unique@example.com",
		IsDefault:    Yes,
	}

	_, err = store.EmailCreate(email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	found := store.EmailFindByEmailAddress("unique@example.com")
	if found == nil {
		t.Fatal("expected to find email by address")
	}

	if found.OwnerID != person.ID {
		t.Errorf("expected OwnerID to be '%s', got '%s'", person.ID, found.OwnerID)
	}
}

func TestEmailFindByEmailAddressNotFound(t *testing.T) {
	store := newTestStore(t)

	found := store.EmailFindByEmailAddress("nonexistent@example.com")
	if found != nil {
		t.Error("expected nil for nonexistent email address")
	}
}

func TestEmailUpdate(t *testing.T) {
	store := newTestStore(t)

	person := &Person{
		Status:    PersonStatusActive,
		FirstName: "Email",
		LastName:  "Updater",
	}

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := &Email{
		OwnerID:      person.ID,
		Status:       EmailStatusActive,
		EmailAddress: "old@example.com",
		IsDefault:    Yes,
	}

	_, err = store.EmailCreate(email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	email.EmailAddress = "new@example.com"
	result := store.EmailUpdate(email)
	if !result {
		t.Error("expected email update to succeed")
	}

	found, err := store.EmailFindByID(email.ID)
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}

	if found.EmailAddress != "new@example.com" {
		t.Errorf("expected EmailAddress to be 'new@example.com', got '%s'", found.EmailAddress)
	}
}

func TestEmailDeleteByID(t *testing.T) {
	store := newTestStore(t)

	person := &Person{
		Status:    PersonStatusActive,
		FirstName: "Email",
		LastName:  "Deleter",
	}

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := &Email{
		OwnerID:      person.ID,
		Status:       EmailStatusActive,
		EmailAddress: "delete@example.com",
		IsDefault:    Yes,
	}

	_, err = store.EmailCreate(email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	deleted, err := store.EmailDeleteByID(email.ID)
	if err != nil {
		t.Fatalf("failed to delete email: %v", err)
	}

	if !deleted {
		t.Error("expected email to be deleted")
	}

	found, err := store.EmailFindByID(email.ID)
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}

	if found != nil {
		t.Error("expected nil email after delete")
	}
}

func TestEmailList(t *testing.T) {
	store := newTestStore(t)

	person := &Person{
		Status:    PersonStatusActive,
		FirstName: "Email",
		LastName:  "Lister",
	}

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	emails := []*Email{
		{OwnerID: person.ID, Status: EmailStatusActive, EmailAddress: "email1@example.com", IsDefault: Yes},
		{OwnerID: person.ID, Status: EmailStatusActive, EmailAddress: "email2@example.com", IsDefault: No},
		{OwnerID: person.ID, Status: EmailStatusActive, EmailAddress: "email3@example.com", IsDefault: No},
	}

	for _, e := range emails {
		_, err = store.EmailCreate(e)
		if err != nil {
			t.Fatalf("failed to create email: %v", err)
		}
	}

	list, err := store.EmailList(&EmailQueryOptions{
		OwnerID: person.ID,
	})
	if err != nil {
		t.Fatalf("failed to list emails: %v", err)
	}

	if len(list) != 3 {
		t.Errorf("expected 3 emails, got %d", len(list))
	}
}
