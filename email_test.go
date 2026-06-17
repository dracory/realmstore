package realmstore

import (
	"context"
	"testing"
)

func TestEmailCreate(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Owner")

	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetStatus(EmailStatusActive)
	email.SetEmailAddress("test@example.com")
	email.SetIsDefault(Yes)

	err = store.EmailCreate(context.Background(), email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	if email.GetID() == "" {
		t.Error("expected email ID to be set")
	}
}

func TestEmailCreateNil(t *testing.T) {
	store := newTestStore(t)

	err := store.EmailCreate(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error for nil email")
	}
}

func TestEmailFindByID(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Finder")

	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetStatus(EmailStatusActive)
	email.SetEmailAddress("findme@example.com")
	email.SetIsDefault(Yes)

	err = store.EmailCreate(context.Background(), email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	found, err := store.EmailFindByID(context.Background(), email.GetID())
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

func TestEmailFindByIDEmpty(t *testing.T) {
	store := newTestStore(t)

	_, err := store.EmailFindByID(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty email ID")
	}
}

func TestEmailFindByEmailAddress(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Address")

	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetStatus(EmailStatusActive)
	email.SetEmailAddress("unique@example.com")
	email.SetIsDefault(Yes)

	err = store.EmailCreate(context.Background(), email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	found, err := store.EmailFindByEmailAddress(context.Background(), "unique@example.com")
	if err != nil {
		t.Fatalf("failed to find email by address: %v", err)
	}
	if found == nil {
		t.Fatal("expected to find email by address")
	}

	if found.GetOwnerID() != person.GetID() {
		t.Errorf("expected OwnerID to be '%s', got '%s'", person.GetID(), found.GetOwnerID())
	}
}

func TestEmailFindByEmailAddressEmpty(t *testing.T) {
	store := newTestStore(t)

	_, err := store.EmailFindByEmailAddress(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty email address")
	}
}

func TestEmailFindByEmailAddressNotFound(t *testing.T) {
	store := newTestStore(t)

	found, err := store.EmailFindByEmailAddress(context.Background(), "nonexistent@example.com")
	if err != nil {
		t.Fatalf("failed to find email by address: %v", err)
	}
	if found != nil {
		t.Error("expected nil for nonexistent email address")
	}
}

func TestEmailFindByPersonID(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetFirstName("Owner")
	store.PersonCreate(context.Background(), person)

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetEmailAddress("owner@example.com")
	store.EmailCreate(context.Background(), email)

	found, err := store.EmailFindByPersonID(context.Background(), person.GetID())
	if err != nil {
		t.Fatalf("failed to find email by person ID: %v", err)
	}
	if found == nil {
		t.Fatal("expected to find email by person ID")
	}
	if found.GetEmailAddress() != "owner@example.com" {
		t.Errorf("expected email 'owner@example.com', got '%s'", found.GetEmailAddress())
	}
}

func TestEmailFindByPersonIDEmpty(t *testing.T) {
	store := newTestStore(t)

	_, err := store.EmailFindByPersonID(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty person ID")
	}
}

func TestEmailFind(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetFirstName("Finder")
	store.PersonCreate(context.Background(), person)

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetEmailAddress("find@example.com")
	email.SetStatus(EmailStatusActive)
	store.EmailCreate(context.Background(), email)

	found, err := store.EmailFind(context.Background(), EmailQuery().
		SetOwnerID(person.GetID()).
		SetStatus(EmailStatusActive))
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}
	if found == nil {
		t.Fatal("expected to find email")
	}
	if found.GetEmailAddress() != "find@example.com" {
		t.Errorf("expected email 'find@example.com', got '%s'", found.GetEmailAddress())
	}
}

func TestEmailUpdate(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Updater")

	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetStatus(EmailStatusActive)
	email.SetEmailAddress("old@example.com")
	email.SetIsDefault(Yes)

	err = store.EmailCreate(context.Background(), email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	email.SetEmailAddress("new@example.com")
	err = store.EmailUpdate(context.Background(), email)
	if err != nil {
		t.Fatalf("failed to update email: %v", err)
	}

	found, err := store.EmailFindByID(context.Background(), email.GetID())
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}

	if found.GetEmailAddress() != "new@example.com" {
		t.Errorf("expected EmailAddress to be 'new@example.com', got '%s'", found.GetEmailAddress())
	}
}

func TestEmailUpdateNil(t *testing.T) {
	store := newTestStore(t)

	err := store.EmailUpdate(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error for nil email")
	}
}

func TestEmailUpdateByID(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetFirstName("Updater")
	store.PersonCreate(context.Background(), person)

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetEmailAddress("old@example.com")
	store.EmailCreate(context.Background(), email)

	err := store.EmailUpdateByID(context.Background(), email.GetID(), map[string]any{
		"email_address": "new@example.com",
	})
	if err != nil {
		t.Fatalf("failed to update email by ID: %v", err)
	}

	found, err := store.EmailFindByID(context.Background(), email.GetID())
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}
	if found.GetEmailAddress() != "new@example.com" {
		t.Errorf("expected email 'new@example.com', got '%s'", found.GetEmailAddress())
	}
}

func TestEmailUpdateByIDEmpty(t *testing.T) {
	store := newTestStore(t)

	err := store.EmailUpdateByID(context.Background(), "", map[string]any{
		"email_address": "test@example.com",
	})
	if err == nil {
		t.Fatal("expected error for empty email ID")
	}
}

func TestEmailDeleteByID(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Deleter")

	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetStatus(EmailStatusActive)
	email.SetEmailAddress("delete@example.com")
	email.SetIsDefault(Yes)

	err = store.EmailCreate(context.Background(), email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	err = store.EmailDeleteByID(context.Background(), email.GetID())
	if err != nil {
		t.Fatalf("failed to delete email: %v", err)
	}

	found, err := store.EmailFindByID(context.Background(), email.GetID())
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}

	if found != nil {
		t.Error("expected nil email after delete")
	}
}

func TestEmailDeleteByIDEmpty(t *testing.T) {
	store := newTestStore(t)

	err := store.EmailDeleteByID(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty email ID")
	}
}

func TestEmailList(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Email")
	person.SetLastName("Lister")

	err := store.PersonCreate(context.Background(), person)
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
		err = store.EmailCreate(context.Background(), e)
		if err != nil {
			t.Fatalf("failed to create email: %v", err)
		}
	}

	list, err := store.EmailList(context.Background(), EmailQuery().SetOwnerID(person.GetID()))
	if err != nil {
		t.Fatalf("failed to list emails: %v", err)
	}

	if len(list) != 3 {
		t.Errorf("expected 3 emails, got %d", len(list))
	}
}

func TestEmailListWithLimitOffset(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	store.PersonCreate(context.Background(), person)

	for i := 0; i < 5; i++ {
		email := NewEmail()
		email.SetOwnerID(person.GetID())
		email.SetEmailAddress("email" + string(rune('1'+i)) + "@example.com")
		store.EmailCreate(context.Background(), email)
	}

	list, err := store.EmailList(context.Background(), EmailQuery().
		SetOwnerID(person.GetID()).
		SetLimit(2).
		SetOffset(1))
	if err != nil {
		t.Fatalf("failed to list emails: %v", err)
	}
	if len(list) != 2 {
		t.Errorf("expected 2 emails, got %d", len(list))
	}
}

func TestEmailListWithStatusFilter(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	store.PersonCreate(context.Background(), person)

	active := NewEmail()
	active.SetOwnerID(person.GetID())
	active.SetStatus(EmailStatusActive)
	active.SetEmailAddress("active@example.com")
	store.EmailCreate(context.Background(), active)

	inactive := NewEmail()
	inactive.SetOwnerID(person.GetID())
	inactive.SetStatus(EmailStatusInactive)
	inactive.SetEmailAddress("inactive@example.com")
	store.EmailCreate(context.Background(), inactive)

	list, err := store.EmailList(context.Background(), EmailQuery().
		SetOwnerID(person.GetID()).
		SetStatus(EmailStatusActive))
	if err != nil {
		t.Fatalf("failed to list emails: %v", err)
	}
	if len(list) != 1 {
		t.Errorf("expected 1 active email, got %d", len(list))
	}
	if list[0].GetEmailAddress() != "active@example.com" {
		t.Errorf("expected 'active@example.com', got '%s'", list[0].GetEmailAddress())
	}
}

func TestEmailListWithIsDefaultFilter(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	store.PersonCreate(context.Background(), person)

	def := NewEmail()
	def.SetOwnerID(person.GetID())
	def.SetIsDefault(Yes)
	def.SetEmailAddress("default@example.com")
	store.EmailCreate(context.Background(), def)

	other := NewEmail()
	other.SetOwnerID(person.GetID())
	other.SetIsDefault(No)
	other.SetEmailAddress("other@example.com")
	store.EmailCreate(context.Background(), other)

	list, err := store.EmailList(context.Background(), EmailQuery().
		SetOwnerID(person.GetID()).
		SetIsDefault(Yes))
	if err != nil {
		t.Fatalf("failed to list emails: %v", err)
	}
	if len(list) != 1 {
		t.Errorf("expected 1 default email, got %d", len(list))
	}
	if list[0].GetEmailAddress() != "default@example.com" {
		t.Errorf("expected 'default@example.com', got '%s'", list[0].GetEmailAddress())
	}
}

func TestEmailListWithSearch(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	store.PersonCreate(context.Background(), person)

	john := NewEmail()
	john.SetOwnerID(person.GetID())
	john.SetEmailAddress("john@example.com")
	store.EmailCreate(context.Background(), john)

	jane := NewEmail()
	jane.SetOwnerID(person.GetID())
	jane.SetEmailAddress("jane@example.com")
	store.EmailCreate(context.Background(), jane)

	list, err := store.EmailList(context.Background(), EmailQuery().
		SetOwnerID(person.GetID()).
		SetSearch("jane"))
	if err != nil {
		t.Fatalf("failed to list emails: %v", err)
	}
	if len(list) != 1 {
		t.Errorf("expected 1 email, got %d", len(list))
	}
	if list[0].GetEmailAddress() != "jane@example.com" {
		t.Errorf("expected 'jane@example.com', got '%s'", list[0].GetEmailAddress())
	}
}

func TestEmailListWithSort(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	store.PersonCreate(context.Background(), person)

	for i := 0; i < 3; i++ {
		email := NewEmail()
		email.SetOwnerID(person.GetID())
		email.SetEmailAddress(string(rune('c'-i)) + "@example.com")
		store.EmailCreate(context.Background(), email)
	}

	list, err := store.EmailList(context.Background(), EmailQuery().
		SetOwnerID(person.GetID()).
		SetSortBy("email_address").
		SetSortOrder("asc").
		SetLimit(10))
	if err != nil {
		t.Fatalf("failed to list emails: %v", err)
	}
	if len(list) < 3 {
		t.Fatalf("expected at least 3 emails, got %d", len(list))
	}
	if list[0].GetEmailAddress() > list[1].GetEmailAddress() {
		t.Error("expected ascending sort")
	}
}

func TestEmailCount(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	store.PersonCreate(context.Background(), person)

	for i := 0; i < 4; i++ {
		email := NewEmail()
		email.SetOwnerID(person.GetID())
		email.SetEmailAddress("email" + string(rune('1'+i)) + "@example.com")
		store.EmailCreate(context.Background(), email)
	}

	count, err := store.EmailCount(context.Background(), EmailQuery().SetOwnerID(person.GetID()))
	if err != nil {
		t.Fatalf("failed to count emails: %v", err)
	}
	if count != 4 {
		t.Errorf("expected count 4, got %d", count)
	}
}

func TestEmailSoftDeleteExclusion(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	store.PersonCreate(context.Background(), person)

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetEmailAddress("delete@example.com")
	store.EmailCreate(context.Background(), email)

	store.EmailDeleteByID(context.Background(), email.GetID())

	list, err := store.EmailList(context.Background(), EmailQuery().SetWithSoftDeleted(true))
	if err != nil {
		t.Fatalf("failed to list emails: %v", err)
	}
	found := false
	for _, e := range list {
		if e.GetID() == email.GetID() {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find soft-deleted email when WithSoftDeleted is true")
	}

	list, err = store.EmailList(context.Background(), EmailQuery().SetDeletedOnly(true))
	if err != nil {
		t.Fatalf("failed to list emails: %v", err)
	}
	found = false
	for _, e := range list {
		if e.GetID() == email.GetID() {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find soft-deleted email when DeletedOnly is true")
	}
}

func TestEmailEntityGettersSetters(t *testing.T) {
	email := NewEmail()

	if email.GetID() == "" {
		t.Error("expected non-empty ID")
	}

	email.SetID("e1")
	if email.GetID() != "e1" {
		t.Errorf("expected ID 'e1', got '%s'", email.GetID())
	}

	email.SetOwnerID("p1")
	if email.GetOwnerID() != "p1" {
		t.Errorf("expected owner ID 'p1', got '%s'", email.GetOwnerID())
	}

	email.SetStatus(EmailStatusInactive)
	if email.GetStatus() != EmailStatusInactive {
		t.Errorf("expected status '%s', got '%s'", EmailStatusInactive, email.GetStatus())
	}

	email.SetEmailAddress("test@example.com")
	if email.GetEmailAddress() != "test@example.com" {
		t.Errorf("expected email 'test@example.com', got '%s'", email.GetEmailAddress())
	}

	email.SetIsDefault(Yes)
	if email.GetIsDefault() != Yes {
		t.Errorf("expected is_default '%s', got '%s'", Yes, email.GetIsDefault())
	}

	email.SetCreatedAt("2024-01-01 12:00:00")
	if email.GetCreatedAt() == "" {
		t.Error("expected created at to be set")
	}
	if email.GetCreatedAtCarbon() == nil {
		t.Error("expected created at carbon to be non-nil")
	}

	email.SetUpdatedAt("2024-01-02 12:00:00")
	if email.GetUpdatedAt() == "" {
		t.Error("expected updated at to be set")
	}
	if email.GetUpdatedAtCarbon() == nil {
		t.Error("expected updated at carbon to be non-nil")
	}

	email.SetSoftDeletedAt("9999-12-31 23:59:59")
	if email.GetSoftDeletedAt() == "" {
		t.Error("expected soft deleted at to be set")
	}
	if email.GetSoftDeletedAtCarbon() == nil {
		t.Error("expected soft deleted at carbon to be non-nil")
	}
	if email.IsSoftDeleted() {
		t.Error("expected email to not be soft deleted")
	}

	email.SetSoftDeletedAt("2020-01-01 00:00:00")
	if !email.IsSoftDeleted() {
		t.Error("expected email to be soft deleted")
	}

	email.SetSoftDeletedAt("")
	if email.GetSoftDeletedAt() != "" {
		t.Error("expected soft deleted at to be empty after clearing")
	}
}

func TestEmailFromExistingData(t *testing.T) {
	data := map[string]string{
		"id":              "e1",
		"owner_id":        "p1",
		"status":          EmailStatusActive,
		"email_address":   "test@example.com",
		"is_default":      Yes,
		"created_at":      "2024-01-01 12:00:00",
		"updated_at":      "2024-01-02 12:00:00",
		"soft_deleted_at": "9999-12-31 23:59:59",
	}

	email := NewEmailFromExistingData(data)
	if email.GetID() != "e1" {
		t.Errorf("expected ID 'e1', got '%s'", email.GetID())
	}
	if email.GetOwnerID() != "p1" {
		t.Errorf("expected owner ID 'p1', got '%s'", email.GetOwnerID())
	}
	if email.GetEmailAddress() != "test@example.com" {
		t.Errorf("expected email 'test@example.com', got '%s'", email.GetEmailAddress())
	}
	if email.GetIsDefault() != Yes {
		t.Errorf("expected is_default '%s', got '%s'", Yes, email.GetIsDefault())
	}
}
