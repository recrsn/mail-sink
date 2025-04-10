package email

import (
	"bytes"
	"strings"
	"testing"
)

func TestGetHeaderValue(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string][]string
		key     string
		want    string
	}{
		{
			name:    "existing header",
			headers: map[string][]string{"Content-Type": {"text/plain"}},
			key:     "Content-Type",
			want:    "text/plain",
		},
		{
			name:    "non-existent header",
			headers: map[string][]string{"Content-Type": {"text/plain"}},
			key:     "Subject",
			want:    "",
		},
		{
			name:    "multiple values, return first",
			headers: map[string][]string{"To": {"user1@example.com", "user2@example.com"}},
			key:     "To",
			want:    "user1@example.com",
		},
		{
			name:    "empty array",
			headers: map[string][]string{"Empty": {}},
			key:     "Empty",
			want:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getHeaderValue(tt.headers, tt.key)
			if got != tt.want {
				t.Errorf("getHeaderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractMessageParts_SimpleText(t *testing.T) {
	// Test a simple text email
	msg := "This is a simple text email"
	headers := map[string][]string{
		"Content-Type": {"text/plain"},
	}

	textBody, htmlBody, attachments := ExtractMessageParts(strings.NewReader(msg), headers)

	if textBody != msg {
		t.Errorf("Expected text body to be %q, got %q", msg, textBody)
	}

	if htmlBody != "" {
		t.Errorf("Expected HTML body to be empty, got %q", htmlBody)
	}

	if len(attachments) != 0 {
		t.Errorf("Expected no attachments, got %d", len(attachments))
	}
}

func TestExtractMessageParts_MultipartMessage(t *testing.T) {
	// Create a multipart message with text, HTML, and an attachment
	boundary := "boundary123"
	
	var body bytes.Buffer
	body.WriteString("--" + boundary + "\r\n")
	body.WriteString("Content-Type: text/plain\r\n\r\n")
	body.WriteString("This is the text part\r\n")
	body.WriteString("--" + boundary + "\r\n")
	body.WriteString("Content-Type: text/html\r\n\r\n")
	body.WriteString("<p>This is the HTML part</p>\r\n")
	body.WriteString("--" + boundary + "\r\n")
	body.WriteString("Content-Type: application/pdf\r\n")
	body.WriteString("Content-Disposition: attachment; filename=\"test.pdf\"\r\n\r\n")
	body.WriteString("PDF content here\r\n")
	body.WriteString("--" + boundary + "--\r\n")

	headers := map[string][]string{
		"Content-Type": {"multipart/mixed; boundary=" + boundary},
	}

	textBody, htmlBody, attachments := ExtractMessageParts(&body, headers)

	if textBody != "This is the text part" {
		t.Errorf("Expected text body to be %q, got %q", "This is the text part", textBody)
	}

	if htmlBody != "<p>This is the HTML part</p>" {
		t.Errorf("Expected HTML body to be %q, got %q", "<p>This is the HTML part</p>", htmlBody)
	}

	if len(attachments) != 1 {
		t.Errorf("Expected 1 attachment, got %d", len(attachments))
	} else {
		att := attachments[0]
		if att.Filename != "test.pdf" {
			t.Errorf("Expected attachment filename to be 'test.pdf', got %q", att.Filename)
		}
		if att.ContentType != "application/pdf" {
			t.Errorf("Expected attachment content type to be 'application/pdf', got %q", att.ContentType)
		}
	}
}