# Makefile for Mail Sink

# Binary name
BINARY_NAME=mail-sink
VERSION=$(shell git describe --tags --always)
COMMIT=$(shell git rev-parse --short HEAD)

# Default target
.PHONY: all
all: $(BINARY_NAME)

# Build Go binary (depends on UI being built)
$(BINARY_NAME): main.go go.mod ui/dist
	@echo "Building Mail Sink..."
	go build -ldflags "-X github.com/recrsn/mail-sink/internal/version.GitCommit=$(COMMIT) -X github.com/recrsn/mail-sink/internal/version.Version=$(VERSION)" -o $(BINARY_NAME)

# UI build - only rebuilds if src files are newer than dist
ui/dist: ui/src/**/* ui/package.json ui/vite.config.ts
	@echo "Building UI..."
	cd ui && npm install && npm run build
	@touch ui/dist

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	rm -rf ui/dist

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

# Help target
.PHONY: help
help:
	@echo "Mail Sink Makefile"
	@echo ""
	@echo "Targets:"
	@echo "  all       - Build the entire application (default)"
	@echo "  clean     - Remove build artifacts"
	@echo "  test      - Run all tests"