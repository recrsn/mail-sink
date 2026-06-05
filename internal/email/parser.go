package email

import (
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"strings"
)

func ExtractMessageParts(msg io.Reader, headers map[string][]string) (textBody, htmlBody string, attachments []Attachment) {
	mediaType, params, err := mime.ParseMediaType(getHeaderValue(headers, "Content-Type"))
	if err != nil || !strings.HasPrefix(mediaType, "multipart/") {
		// Not a multipart message, treat the body as text
		body, _ := io.ReadAll(transferDecodedReader(msg, getHeaderValue(headers, "Content-Transfer-Encoding")))
		return string(body), "", nil
	}

	boundary := params["boundary"]
	if boundary == "" {
		// Invalid multipart message
		body, _ := io.ReadAll(transferDecodedReader(msg, getHeaderValue(headers, "Content-Transfer-Encoding")))
		return string(body), "", nil
	}

	mr := multipart.NewReader(msg, boundary)
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		content, err := io.ReadAll(transferDecodedReader(part, part.Header.Get("Content-Transfer-Encoding")))
		if err != nil {
			continue
		}

		partMediaType, _, err := mime.ParseMediaType(part.Header.Get("Content-Type"))
		if err != nil {
			continue
		}

		contentDisposition := part.Header.Get("Content-Disposition")
		_, dispositionParams, err := mime.ParseMediaType(contentDisposition)
		filename := dispositionParams["filename"]

		if strings.HasPrefix(partMediaType, "text/plain") {
			textBody = string(content)
		} else if strings.HasPrefix(partMediaType, "text/html") {
			htmlBody = string(content)
		} else if filename != "" || strings.HasPrefix(partMediaType, "application/") || strings.HasPrefix(partMediaType, "image/") {
			// It's likely an attachment
			attachments = append(attachments, Attachment{
				Filename:      filename,
				ContentType:   partMediaType,
				Size:          len(content),
				Content:       string(content),
				ContentBase64: content,
			})
		}
	}

	return textBody, htmlBody, attachments
}

func transferDecodedReader(r io.Reader, encoding string) io.Reader {
	switch strings.ToLower(strings.TrimSpace(encoding)) {
	case "quoted-printable":
		return quotedprintable.NewReader(r)
	case "base64":
		return base64.NewDecoder(base64.StdEncoding, r)
	default:
		return r
	}
}

func getHeaderValue(headers map[string][]string, key string) string {
	if values, ok := headers[key]; ok && len(values) > 0 {
		return values[0]
	}
	return ""
}
