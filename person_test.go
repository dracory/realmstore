package realmstore

import (
	"context"
	"testing"
	"time"
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

func TestPersonCreateNil(t *testing.T) {
	store := newTestStore(t)

	err := store.PersonCreate(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error for nil person")
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

func TestPersonFindByIDEmpty(t *testing.T) {
	store := newTestStore(t)

	_, err := store.PersonFindByID(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty person ID")
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

func TestPersonUpdateNil(t *testing.T) {
	store := newTestStore(t)

	err := store.PersonUpdate(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error for nil person")
	}
}

func TestPersonUpdateByID(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetStatus(PersonStatusActive)
	person.SetFirstName("Alice")
	person.SetLastName("Wonder")

	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	err = store.PersonUpdateByID(context.Background(), person.GetID(), map[string]any{
		"first_name": "Alicia",
	})
	if err != nil {
		t.Fatalf("failed to update person by ID: %v", err)
	}

	found, err := store.PersonFindByID(context.Background(), person.GetID())
	if err != nil {
		t.Fatalf("failed to find person: %v", err)
	}

	if found.GetFirstName() != "Alicia" {
		t.Errorf("expected FirstName to be 'Alicia', got '%s'", found.GetFirstName())
	}
}

func TestPersonUpdateByIDEmpty(t *testing.T) {
	store := newTestStore(t)

	err := store.PersonUpdateByID(context.Background(), "", map[string]any{
		"first_name": "Test",
	})
	if err == nil {
		t.Fatal("expected error for empty person ID")
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

func TestPersonDeleteEmpty(t *testing.T) {
	store := newTestStore(t)

	err := store.PersonDelete(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty person ID")
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

func TestPersonCreateWithEmail(t *testing.T) {
	store := newTestStore(t)

	person, err := store.PersonCreateWithEmail(context.Background(), map[string]string{
		"firstName": "Jane",
		"lastName":  "Doe",
		"email":     "jane@example.com",
	})
	if err != nil {
		t.Fatalf("failed to create person with email: %v", err)
	}

	if person.GetID() == "" {
		t.Fatal("expected person ID to be set")
	}

	foundEmail, err := store.EmailFindByEmailAddress(context.Background(), "jane@example.com")
	if err != nil {
		t.Fatalf("failed to find email: %v", err)
	}
	if foundEmail == nil {
		t.Fatal("expected email to be created")
	}
	if foundEmail.GetOwnerID() != person.GetID() {
		t.Errorf("expected email owner to be '%s', got '%s'", person.GetID(), foundEmail.GetOwnerID())
	}
}

func TestPersonFindByEmail(t *testing.T) {
	store := newTestStore(t)

	person, err := store.PersonFindOrCreate(context.Background(), "Alice", "Smith", "alice@example.com")
	if err != nil {
		t.Fatalf("failed to find or create person: %v", err)
	}

	found, err := store.PersonFindByEmail(context.Background(), "alice@example.com")
	if err != nil {
		t.Fatalf("failed to find person by email: %v", err)
	}
	if found == nil {
		t.Fatal("expected to find person by email")
	}
	if found.GetID() != person.GetID() {
		t.Errorf("expected person ID '%s', got '%s'", person.GetID(), found.GetID())
	}

	notFound, err := store.PersonFindByEmail(context.Background(), "nonexistent@example.com")
	if err != nil {
		t.Fatalf("failed to find person by email: %v", err)
	}
	if notFound != nil {
		t.Error("expected nil for nonexistent email")
	}
}

func TestPersonFindDefaultEmail(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetFirstName("Bob")
	person.SetLastName("Builder")
	err := store.PersonCreate(context.Background(), person)
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	email := NewEmail()
	email.SetOwnerID(person.GetID())
	email.SetEmailAddress("bob@example.com")
	email.SetIsDefault(Yes)
	err = store.EmailCreate(context.Background(), email)
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	found, err := store.PersonFindDefaultEmail(context.Background(), person.GetID())
	if err != nil {
		t.Fatalf("failed to find default email: %v", err)
	}
	if found == nil {
		t.Fatal("expected to find default email")
	}
	if found.GetEmailAddress() != "bob@example.com" {
		t.Errorf("expected email 'bob@example.com', got '%s'", found.GetEmailAddress())
	}
}

func TestPersonFindDefaultEmailEmpty(t *testing.T) {
	store := newTestStore(t)

	_, err := store.PersonFindDefaultEmail(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty person ID")
	}
}

func TestPersonList(t *testing.T) {
	store := newTestStore(t)

	for i := 0; i < 5; i++ {
		person := NewPerson()
		person.SetFirstName("Person")
		person.SetLastName(string(rune('A' + i)))
		err := store.PersonCreate(context.Background(), person)
		if err != nil {
			t.Fatalf("failed to create person: %v", err)
		}
	}

	list, err := store.PersonList(context.Background(), PersonQuery().SetLimit(3))
	if err != nil {
		t.Fatalf("failed to list persons: %v", err)
	}
	if len(list) != 3 {
		t.Errorf("expected 3 persons, got %d", len(list))
	}
}

func TestPersonListWithOffset(t *testing.T) {
	store := newTestStore(t)

	for i := 0; i < 5; i++ {
		person := NewPerson()
		person.SetFirstName("Person")
		person.SetLastName(string(rune('A' + i)))
		err := store.PersonCreate(context.Background(), person)
		if err != nil {
			t.Fatalf("failed to create person: %v", err)
		}
	}

	list, err := store.PersonList(context.Background(), PersonQuery().SetLimit(2).SetOffset(2))
	if err != nil {
		t.Fatalf("failed to list persons: %v", err)
	}
	if len(list) != 2 {
		t.Errorf("expected 2 persons, got %d", len(list))
	}
}

func TestPersonListWithStatusFilter(t *testing.T) {
	store := newTestStore(t)

	active := NewPerson()
	active.SetStatus(PersonStatusActive)
	active.SetFirstName("Active")
	store.PersonCreate(context.Background(), active)

	inactive := NewPerson()
	inactive.SetStatus(PersonStatusInactive)
	inactive.SetFirstName("Inactive")
	store.PersonCreate(context.Background(), inactive)

	list, err := store.PersonList(context.Background(), PersonQuery().SetStatus(PersonStatusActive))
	if err != nil {
		t.Fatalf("failed to list persons: %v", err)
	}

	foundActive := false
	for _, p := range list {
		if p.GetFirstName() == "Inactive" {
			t.Error("expected inactive person to be filtered out")
		}
		if p.GetFirstName() == "Active" {
			foundActive = true
		}
	}
	if !foundActive {
		t.Error("expected active person to be found")
	}
}

func TestPersonListWithSearch(t *testing.T) {
	store := newTestStore(t)

	john := NewPerson()
	john.SetFirstName("John")
	john.SetLastName("Doe")
	store.PersonCreate(context.Background(), john)

	jane := NewPerson()
	jane.SetFirstName("Jane")
	jane.SetLastName("Smith")
	store.PersonCreate(context.Background(), jane)

	list, err := store.PersonList(context.Background(), PersonQuery().SetSearch("Jane"))
	if err != nil {
		t.Fatalf("failed to list persons: %v", err)
	}

	foundJane := false
	for _, p := range list {
		if p.GetFirstName() == "John" {
			t.Error("expected John to be filtered out")
		}
		if p.GetFirstName() == "Jane" {
			foundJane = true
		}
	}
	if !foundJane {
		t.Error("expected Jane to be found")
	}
}

func TestPersonListWithSort(t *testing.T) {
	store := newTestStore(t)

	for i := 0; i < 3; i++ {
		person := NewPerson()
		person.SetFirstName(string(rune('C' - i)))
		store.PersonCreate(context.Background(), person)
	}

	list, err := store.PersonList(context.Background(), PersonQuery().
		SetSortBy("first_name").
		SetSortOrder("asc").
		SetLimit(10))
	if err != nil {
		t.Fatalf("failed to list persons: %v", err)
	}
	if len(list) < 3 {
		t.Fatalf("expected at least 3 persons, got %d", len(list))
	}
	if list[0].GetFirstName() > list[1].GetFirstName() {
		t.Error("expected ascending sort")
	}
}

func TestPersonCount(t *testing.T) {
	store := newTestStore(t)

	for i := 0; i < 4; i++ {
		person := NewPerson()
		person.SetFirstName("Person")
		store.PersonCreate(context.Background(), person)
	}

	count, err := store.PersonCount(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to count persons: %v", err)
	}
	if count != 4 {
		t.Errorf("expected count 4, got %d", count)
	}
}

func TestPersonMaxPages(t *testing.T) {
	store := newTestStore(t)

	for i := 0; i < 25; i++ {
		person := NewPerson()
		person.SetFirstName("Person")
		store.PersonCreate(context.Background(), person)
	}

	pages, err := store.PersonMaxPages(context.Background(), PersonQuery().SetLimit(10))
	if err != nil {
		t.Fatalf("failed to get max pages: %v", err)
	}
	if pages != 3 {
		t.Errorf("expected 3 pages, got %d", pages)
	}
}

func TestPersonMaxPagesExactDivision(t *testing.T) {
	store := newTestStore(t)

	for i := 0; i < 20; i++ {
		person := NewPerson()
		person.SetFirstName("Person")
		store.PersonCreate(context.Background(), person)
	}

	pages, err := store.PersonMaxPages(context.Background(), PersonQuery().SetLimit(10))
	if err != nil {
		t.Fatalf("failed to get max pages: %v", err)
	}
	if pages != 2 {
		t.Errorf("expected 2 pages, got %d", pages)
	}
}

func TestPersonSoftDeleteExclusion(t *testing.T) {
	store := newTestStore(t)

	person := NewPerson()
	person.SetFirstName("SoftDelete")
	store.PersonCreate(context.Background(), person)

	store.PersonDelete(context.Background(), person.GetID())

	list, err := store.PersonList(context.Background(), PersonQuery().SetWithSoftDeleted(true))
	if err != nil {
		t.Fatalf("failed to list persons: %v", err)
	}
	found := false
	for _, p := range list {
		if p.GetID() == person.GetID() {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find soft-deleted person when WithSoftDeleted is true")
	}

	list, err = store.PersonList(context.Background(), PersonQuery().SetDeletedOnly(true))
	if err != nil {
		t.Fatalf("failed to list persons: %v", err)
	}
	found = false
	for _, p := range list {
		if p.GetID() == person.GetID() {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find soft-deleted person when DeletedOnly is true")
	}
}

func TestPersonEntityGettersSetters(t *testing.T) {
	person := NewPerson()

	if person.GetID() == "" {
		t.Error("expected non-empty ID")
	}

	person.SetID("test-id")
	if person.GetID() != "test-id" {
		t.Errorf("expected ID 'test-id', got '%s'", person.GetID())
	}

	person.SetStatus(PersonStatusInactive)
	if person.GetStatus() != PersonStatusInactive {
		t.Errorf("expected status '%s', got '%s'", PersonStatusInactive, person.GetStatus())
	}

	person.SetFirstName("Alice")
	if person.GetFirstName() != "Alice" {
		t.Errorf("expected first name 'Alice', got '%s'", person.GetFirstName())
	}

	person.SetMiddleNames("Marie")
	if person.GetMiddleNames() != "Marie" {
		t.Errorf("expected middle names 'Marie', got '%s'", person.GetMiddleNames())
	}

	person.SetLastName("Wonder")
	if person.GetLastName() != "Wonder" {
		t.Errorf("expected last name 'Wonder', got '%s'", person.GetLastName())
	}

	birth := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	person.SetBirthDay(&birth)
	if person.GetBirthDay() == nil || !person.GetBirthDay().Equal(birth) {
		t.Error("expected birth day to be set")
	}

	person.SetBirthDay(nil)
	if person.GetBirthDay() != nil {
		t.Error("expected birth day to be nil")
	}

	person.SetCreatedAt("2024-01-01 12:00:00")
	if person.GetCreatedAt() == "" {
		t.Error("expected created at to be set")
	}
	if person.GetCreatedAtCarbon() == nil {
		t.Error("expected created at carbon to be non-nil")
	}

	person.SetUpdatedAt("2024-01-02 12:00:00")
	if person.GetUpdatedAt() == "" {
		t.Error("expected updated at to be set")
	}
	if person.GetUpdatedAtCarbon() == nil {
		t.Error("expected updated at carbon to be non-nil")
	}

	person.SetSoftDeletedAt("9999-12-31 23:59:59")
	if person.GetSoftDeletedAt() == "" {
		t.Error("expected soft deleted at to be set")
	}
	if person.GetSoftDeletedAtCarbon() == nil {
		t.Error("expected soft deleted at carbon to be non-nil")
	}
	if person.IsSoftDeleted() {
		t.Error("expected person to not be soft deleted")
	}

	person.SetSoftDeletedAt("2020-01-01 00:00:00")
	if !person.IsSoftDeleted() {
		t.Error("expected person to be soft deleted")
	}

	person.SetSoftDeletedAt("")
	if person.GetSoftDeletedAt() != "" {
		t.Error("expected soft deleted at to be empty after clearing")
	}
}

func TestPersonFromExistingData(t *testing.T) {
	data := map[string]string{
		"id":              "p1",
		"status":          PersonStatusActive,
		"first_name":      "Alice",
		"middle_names":    "Marie",
		"last_name":       "Smith",
		"created_at":      "2024-01-01 12:00:00",
		"updated_at":      "2024-01-02 12:00:00",
		"soft_deleted_at": "9999-12-31 23:59:59",
	}

	person := NewPersonFromExistingData(data)
	if person.GetID() != "p1" {
		t.Errorf("expected ID 'p1', got '%s'", person.GetID())
	}
	if person.GetFirstName() != "Alice" {
		t.Errorf("expected first name 'Alice', got '%s'", person.GetFirstName())
	}
	if person.GetStatus() != PersonStatusActive {
		t.Errorf("expected status '%s', got '%s'", PersonStatusActive, person.GetStatus())
	}
}

func TestPersonWithEmailInterface(t *testing.T) {
	pwe := &personWithEmail{}
	pwe.SetID("p1")
	pwe.SetEmailAddress("test@example.com")

	if pwe.GetEmailAddress() != "test@example.com" {
		t.Errorf("expected email 'test@example.com', got '%s'", pwe.GetEmailAddress())
	}

	var iface PersonWithEmailInterface = pwe
	if iface.GetID() != "p1" {
		t.Errorf("expected ID 'p1', got '%s'", iface.GetID())
	}
}

func TestPersonQueryValidation(t *testing.T) {
	q := NewPersonQuery()
	if err := q.Validate(); err != nil {
		t.Fatalf("expected valid query, got %v", err)
	}

	q.SetID("")
	if err := q.Validate(); err == nil {
		t.Error("expected error for empty ID in PersonQuery")
	}

	q = NewPersonQuery().SetLimit(-1)
	if err := q.Validate(); err == nil {
		t.Error("expected error for negative limit in PersonQuery")
	}

	q = NewPersonQuery().SetOffset(-1)
	if err := q.Validate(); err == nil {
		t.Error("expected error for negative offset in PersonQuery")
	}

	q = NewPersonQuery().SetWithEmail(true)
	if !q.IsWithEmailSet() || !q.GetWithEmail() {
		t.Error("expected WithEmail to be set")
	}
}
