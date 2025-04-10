package email

import (
	"testing"
	"time"
)

func TestNewStore(t *testing.T) {
	store := NewStore()
	if store == nil {
		t.Fatal("NewStore() returned nil")
	}

	if store.emails == nil {
		t.Fatal("store.emails is nil")
	}

	if len(store.emails) != 0 {
		t.Errorf("Expected empty emails slice, got %d elements", len(store.emails))
	}
}

func TestStore_Add(t *testing.T) {
	store := NewStore()
	email := &Email{
		ID:         "test1",
		From:       "sender@example.com",
		To:         []string{"recipient@example.com"},
		Subject:    "Test Subject",
		TextBody:   "Test Body",
		Headers:    map[string][]string{"Content-Type": {"text/plain"}},
		ReceivedAt: time.Now(),
	}

	store.Add(email)

	emails := store.GetAll()
	if len(emails) != 1 {
		t.Fatalf("Expected 1 email, got %d", len(emails))
	}

	if emails[0].ID != "test1" {
		t.Errorf("Expected ID 'test1', got %q", emails[0].ID)
	}

	// Add another email to test prepending
	email2 := &Email{
		ID:         "test2",
		From:       "sender2@example.com",
		To:         []string{"recipient2@example.com"},
		Subject:    "Test Subject 2",
		TextBody:   "Test Body 2",
		Headers:    map[string][]string{"Content-Type": {"text/plain"}},
		ReceivedAt: time.Now(),
	}

	store.Add(email2)

	emails = store.GetAll()
	if len(emails) != 2 {
		t.Fatalf("Expected 2 emails, got %d", len(emails))
	}

	// Most recent email should be first
	if emails[0].ID != "test2" {
		t.Errorf("Expected first email to have ID 'test2', got %q", emails[0].ID)
	}
}

func TestStore_GetByID(t *testing.T) {
	store := NewStore()
	
	// Add test emails
	email1 := &Email{ID: "test1", Subject: "Test 1"}
	email2 := &Email{ID: "test2", Subject: "Test 2"}
	email3 := &Email{ID: "test3", Subject: "Test 3"}
	
	store.Add(email1)
	store.Add(email2)
	store.Add(email3)
	
	// Test finding existing emails
	found := store.GetByID("test2")
	if found == nil {
		t.Fatal("GetByID() returned nil for existing ID")
	}
	if found.ID != "test2" || found.Subject != "Test 2" {
		t.Errorf("GetByID() returned wrong email. Got ID: %q, Subject: %q", found.ID, found.Subject)
	}
	
	// Test non-existent ID
	notFound := store.GetByID("nonexistent")
	if notFound != nil {
		t.Errorf("GetByID() returned non-nil for non-existent ID: %v", notFound)
	}
}

func TestStore_Clear(t *testing.T) {
	store := NewStore()
	
	// Add some test emails
	store.Add(&Email{ID: "test1"})
	store.Add(&Email{ID: "test2"})
	store.Add(&Email{ID: "test3"})
	
	// Verify emails are added
	if len(store.GetAll()) != 3 {
		t.Fatalf("Expected 3 emails before clearing, got %d", len(store.GetAll()))
	}
	
	// Clear the store
	store.Clear()
	
	// Verify emails are removed
	if len(store.GetAll()) != 0 {
		t.Errorf("Expected 0 emails after clearing, got %d", len(store.GetAll()))
	}
}