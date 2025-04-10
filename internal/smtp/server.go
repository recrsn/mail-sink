package smtp

import (
	"fmt"
	"log"
	"time"

	gosmtp "github.com/emersion/go-smtp"
	"github.com/recrsn/mail-sink/internal/email"
)

type Server struct {
	store    *email.Store
	port     int
	server   *gosmtp.Server
}

func NewServer(store *email.Store, port int) *Server {
	backend := NewBackend(store)
	s := gosmtp.NewServer(backend)
	
	s.Addr = fmt.Sprintf(":%d", port)
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024 * 10 // 10MB
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true
	
	return &Server{
		store:  store,
		port:   port,
		server: s,
	}
}

func (s *Server) Start() {
	log.Printf("Starting SMTP server on :%d", s.port)
	
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			log.Fatalf("SMTP server error: %v", err)
		}
	}()
}