package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/recrsn/mail-sink/internal/api"
	"github.com/recrsn/mail-sink/internal/config"
	"github.com/recrsn/mail-sink/internal/email"
	"github.com/recrsn/mail-sink/internal/smtp"
	"github.com/recrsn/mail-sink/internal/version"
)

//go:embed ui/dist
var uiAssets embed.FS

func main() {
	// Parse command line flags
	cfg := config.ParseFlags()
	
	// Display version if requested
	if cfg.Version {
		fmt.Println(version.GetVersion())
		os.Exit(0)
	}
	
	// Create email store
	emailStore := email.NewStore()

	// Start SMTP server
	smtpServer := smtp.NewServer(emailStore, cfg.SMTPPort)
	smtpServer.Start()

	// Start HTTP server
	apiServer := api.NewServer(emailStore, cfg.SMTPPort, cfg.HTTPPort, uiAssets)
	if err := apiServer.Start(); err != nil {
		log.Fatal(err)
	}
}