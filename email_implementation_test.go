package realmstore

import (
	"testing"
)

func TestEmailCreate(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Owner")

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetStatus(EmailStatusActive)
	email.SetEmailAddress("test@example.com")
	email.SetIsDefault(Yes)

	created, err := store.EmailCreate(email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	if !created {
		t.Error("expected email to be created")
	}

	if email.GetID() == "" {
		t.Error("expected email ID to be set")
	}
}

func TestEmailFindByID(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Finder")

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetStatus(EmailStatusActive)
	email.SetEmailAddress("findme@example.com")
	email.SetIsDefault(Yes)

	_, err = store.EmailCreate(email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	found, err := store.EmailFindByID(email.GetID())
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}

	if found == nil {
		t.Fatal("expected to find email")
	}

	if found.GetEmailAddress() != "findme@example.com" {
		t.Errorf("expected EmailAddress to be 'findme@example.com', got '%s'", found.GetEmailAddress())
	}
}

func TestEmailFindByEmailAddress(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Address")

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetStatus(EmailStatusActive)
	email.SetEmailAddress("unique@example.com")
	email.SetIsDefault(Yes)

	_, err = store.EmailCreate(email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	found := store.EmailFindByEmailAddress("unique@example.com")
	if found == nil {
		t.Fatal("expected to find email by address")
	}

	if found.GetOwnerID() != person.GetID() {
		t.Errorf("expected OwnerID to be '%s', got '%s'", person.GetID(), found.GetOwnerID())
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

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Updater")

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetStatus(EmailStatusActive)
	email.SetEmailAddress("old@example.com")
	email.SetIsDefault(Yes)

	_, err = store.EmailCreate(email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	email.SetEmailAddress("new@example.com")
	result := store.EmailUpdate(email)
	if !result {
		t.Error("expected email update to succeed")
	}

	found, err := store.EmailFindByID(email.GetID())
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}

	if found.GetEmailAddress() != "new@example.com" {
		t.Errorf("expected EmailAddress to be 'new@example.com', got '%s'", found.GetEmailAddress())
	}
}

func TestEmailDeleteByID(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Deleter")

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetStatus(EmailStatusActive)
	email.SetEmailAddress("delete@example.com")
	email.SetIsDefault(Yes)

	_, err = store.EmailCreate(email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	deleted, err := store.EmailDeleteByID(email.GetID())
	if err != nil {
		t.Fatalf("failed to delete email: %v", err)
	}

	if !deleted {
		t.Error("expected email to be deleted")
	}

	found, err := store.EmailFindByID(email.GetID())
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}

	if found != nil {
		t.Error("expected nil email after delete")
	}
}

func TestEmailList(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Lister")

	err := store.PersonCreate(person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	emails := []EmailInterface{
		NewEmail(),
		NewEmail(),
		NewEmail(),
	}

	emails[0].SetOwnerID(person.GetID())
	emails[0].SetStatus(EmailStatusActive)
	emails[0].SetEmailAddress("email1@example.com")
	emails[0].SetIsDefault(Yes)

	emails[1].SetOwnerID(person.GetID())
	emails[1].SetStatus(EmailStatusActive)
	emails[1].SetEmailAddress("email2@example.com")
	emails[1].SetIsDefault(No)

	emails[2].SetOwnerID(person.GetID())
	emails[2].SetStatus(EmailStatusActive)
	emails[2].SetEmailAddress("email3@example.com")
	emails[2].SetIsDefault(No)

	for _, e := range emails {
		_, err = store.EmailCreate(e)
		if err != nil {
			t.Fatalf("failed to create email: %v", err)
		}
	}

	list, err := store.EmailList(&EmailQueryOptions{
		OwnerID: person.GetID(),
	})
	if err != nil {
		t.Fatalf("failed to list emails: %v", err)
	}

	if len(list) != 3 {
		t.Errorf("expected 3 emails, got %d", len(list))
	}
}
