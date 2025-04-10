package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	SMTPPort int
	HTTPPort int
	Version  bool
}

func ParseFlags() *Config {
	config := &Config{
		SMTPPort: 1025,
		HTTPPort: 8080,
	}

	// Check environment variables first
	if smtpPortEnv := os.Getenv("SMTP_PORT"); smtpPortEnv != "" {
		if port, err := strconv.Atoi(smtpPortEnv); err == nil {
			config.SMTPPort = port
		}
	}

	if httpPortEnv := os.Getenv("HTTP_PORT"); httpPortEnv != "" {
		if port, err := strconv.Atoi(httpPortEnv); err == nil {
			config.HTTPPort = port
		}
	}

	// Command line flags override environment variables
	flag.IntVar(&config.SMTPPort, "smtp-port", config.SMTPPort, "SMTP server port (env: SMTP_PORT)")
	flag.IntVar(&config.HTTPPort, "http-port", config.HTTPPort, "HTTP server port (env: HTTP_PORT)")
	flag.BoolVar(&config.Version, "version", false, "Show version information")

	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		fmt.Printf("Mail Sink - A testing SMTP server\n\n")
		fmt.Printf("Usage: mail-sink [options]\n\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	return config
}
