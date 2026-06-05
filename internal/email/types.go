package email

import (
	"sync"
	"time"
)

type Attachment struct {
	Filename      string `json:"filename"`
	ContentType   string `json:"contentType"`
	Size          int    `json:"size"`
	Content       string `json:"content"`
	ContentBase64 []byte `json:"content_base64,omitempty"`
}

type Email struct {
	ID          string              `json:"id"`
	From        string              `json:"from"`
	To          []string            `json:"to"`
	Subject     string              `json:"subject"`
	TextBody    string              `json:"textBody,omitempty"`
	HTMLBody    string              `json:"htmlBody,omitempty"`
	Attachments []Attachment        `json:"attachments,omitempty"`
	Headers     map[string][]string `json:"headers"`
	ReceivedAt  time.Time           `json:"receivedAt"`
}

type Store struct {
	mutex  sync.RWMutex
	emails []*Email
}

func NewStore() *Store {
	return &Store{
		emails: make([]*Email, 0),
	}
}

func (s *Store) Add(email *Email) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.emails = append([]*Email{email}, s.emails...)
}

func (s *Store) GetAll() []*Email {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Return a copy to avoid race conditions
	emailsCopy := make([]*Email, len(s.emails))
	copy(emailsCopy, s.emails)
	return emailsCopy
}

func (s *Store) GetByID(id string) *Email {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, email := range s.emails {
		if email.ID == id {
			return email
		}
	}
	return nil
}

func (s *Store) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.emails = make([]*Email, 0)
}
