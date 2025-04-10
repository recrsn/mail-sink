package smtp

import (
	"strings"
	"testing"

	"github.com/recrsn/mail-sink/internal/email"
)

func TestNewBackend(t *testing.T) {
	store := email.NewStore()
	backend := NewBackend(store)

	if backend == nil {
		t.Fatal("NewBackend() returned nil")
	}

	if backend.store != store {
		t.Errorf("backend.store = %p, want %p", backend.store, store)
	}
}

func TestBackend_NewSession(t *testing.T) {
	store := email.NewStore()
	backend := NewBackend(store)
	session, err := backend.NewSession(nil)

	if err != nil {
		t.Fatalf("backend.NewSession() returned error: %v", err)
	}

	if session == nil {
		t.Fatal("backend.NewSession() returned nil")
	}
}

func TestSession_Mail(t *testing.T) {
	store := email.NewStore()
	backend := NewBackend(store)
	session, _ := backend.NewSession(nil)

	err := session.Mail("sender@example.com", nil)
	if err != nil {
		t.Errorf("session.Mail() returned error: %v", err)
	}

	// Check that from address was set
	s := session.(*Session)
	if s.from != "sender@example.com" {
		t.Errorf("session.from = %q, want %q", s.from, "sender@example.com")
	}
}

func TestSession_Rcpt(t *testing.T) {
	store := email.NewStore()
	backend := NewBackend(store)
	session, _ := backend.NewSession(nil)

	// Set up a from address first
	session.Mail("sender@example.com", nil)

	err := session.Rcpt("recipient@example.com", nil)
	if err != nil {
		t.Errorf("session.Rcpt() returned error: %v", err)
	}

	// Check that the recipient was added
	s := session.(*Session)
	if len(s.recipients) != 1 || s.recipients[0] != "recipient@example.com" {
		t.Errorf("session.recipients = %v, want %v", s.recipients, []string{"recipient@example.com"})
	}

	// Add another recipient
	err = session.Rcpt("another@example.com", nil)
	if err != nil {
		t.Errorf("session.Rcpt() second call returned error: %v", err)
	}

	// Check that the second recipient was added
	if len(s.recipients) != 2 || s.recipients[1] != "another@example.com" {
		t.Errorf("session.recipients = %v, want %v", s.recipients, []string{"recipient@example.com", "another@example.com"})
	}
}

func TestSession_Data(t *testing.T) {
	store := email.NewStore()
	backend := NewBackend(store)
	session, _ := backend.NewSession(nil)

	// Set up from and to addresses
	session.Mail("sender@example.com", nil)
	session.Rcpt("recipient@example.com", nil)

	// Create a simple message
	messageContent := "Subject: Test Subject\r\n\r\nThis is the message body."
	reader := strings.NewReader(messageContent)

	// Call Data with the reader
	err := session.Data(reader)
	if err != nil {
		t.Fatalf("session.Data() returned error: %v", err)
	}

	// Verify that the email was stored
	emails := store.GetAll()
	if len(emails) != 1 {
		t.Fatalf("Expected 1 email in store, got %d", len(emails))
	}

	// Check email fields
	storedEmail := emails[0]
	if storedEmail.From != "sender@example.com" {
		t.Errorf("storedEmail.From = %q, want %q", storedEmail.From, "sender@example.com")
	}

	if len(storedEmail.To) != 1 || storedEmail.To[0] != "recipient@example.com" {
		t.Errorf("storedEmail.To = %v, want %v", storedEmail.To, []string{"recipient@example.com"})
	}

	if storedEmail.Subject != "Test Subject" {
		t.Errorf("storedEmail.Subject = %q, want %q", storedEmail.Subject, "Test Subject")
	}

	if !strings.Contains(storedEmail.TextBody, "This is the message body") {
		t.Errorf("storedEmail.TextBody = %q, should contain %q", storedEmail.TextBody, "This is the message body")
	}
}

func TestSession_Reset(t *testing.T) {
	store := email.NewStore()
	backend := NewBackend(store)
	session, _ := backend.NewSession(nil)
	s := session.(*Session)

	// Set up some state
	session.Mail("sender@example.com", nil)
	session.Rcpt("recipient@example.com", nil)

	// Reset the session
	session.Reset()

	// Verify state was reset
	if s.from != "" {
		t.Errorf("after Reset(), session.from = %q, want empty string", s.from)
	}

	if len(s.recipients) != 0 {
		t.Errorf("after Reset(), session.recipients = %v, want empty slice", s.recipients)
	}
}

func TestSession_Logout(t *testing.T) {
	store := email.NewStore()
	backend := NewBackend(store)
	session, _ := backend.NewSession(nil)

	// Logout should not return an error
	err := session.Logout()
	if err != nil {
		t.Errorf("session.Logout() returned error: %v", err)
	}
}
