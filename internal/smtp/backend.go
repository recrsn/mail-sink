package smtp

import (
	"fmt"
	"io"
	"net/mail"
	"strings"
	"sync/atomic"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/recrsn/mail-sink/internal/email"
)

var emailCounter int64

type Backend struct {
	store *email.Store
}

func NewBackend(store *email.Store) *Backend {
	return &Backend{store: store}
}

func (b *Backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	return &Session{store: b.store}, nil
}

type Session struct {
	store      *email.Store
	from       string
	recipients []string
}

func (s *Session) AuthPlain(username, password string) error {
	// Allow all authentication attempts
	return nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	s.from = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	s.recipients = append(s.recipients, to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	buf, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	msg, err := mail.ReadMessage(strings.NewReader(string(buf)))
	if err != nil {
		return err
	}

	// Convert mail.Header to map[string][]string
	headers := make(map[string][]string)
	for k, vs := range msg.Header {
		headers[k] = vs
	}

	textBody, htmlBody, attachments := email.ExtractMessageParts(msg.Body, headers)

	now := time.Now()
	id := fmt.Sprintf("%d-%d", now.UnixNano(), atomic.AddInt64(&emailCounter, 1))

	email := &email.Email{
		ID:          id,
		From:        s.from,
		To:          s.recipients,
		Subject:     msg.Header.Get("Subject"),
		TextBody:    textBody,
		HTMLBody:    htmlBody,
		Attachments: attachments,
		Headers:     headers,
		ReceivedAt:  now,
	}

	s.store.Add(email)
	return nil
}

func (s *Session) Reset() {
	s.from = ""
	s.recipients = []string{}
}

func (s *Session) Logout() error {
	return nil
}