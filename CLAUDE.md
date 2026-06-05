# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run Commands
- Backend build: `make` or `make all`
- Backend run: `make run` (default ports) or `make run-custom SMTP_PORT=2525 HTTP_PORT=3000`
- UI dev server: `cd ui && npm run dev`
- UI build: `cd ui && npm run build`
- UI lint: `cd ui && npm run lint`

## Code Style Guidelines
- **Go**: Use PascalCase for exported identifiers, camelCase for unexported
- **React**: Use PascalCase for components, camelCase for functions/variables
- **Error Handling**: In Go, follow the standard `if err != nil` pattern with early returns
- **TypeScript**: Use strong typing with interfaces for component props and API responses
- **Components**: Follow shadcn/ui patterns for new components, use the shadcn CLI to add components
- **Package Structure**: Maintain domain-based organization (`internal/api`, `internal/smtp`, etc.)
- **Naming**: Use descriptive names that reflect purpose, avoid abbreviations
- **File Naming**: snake_case for Go files, PascalCase.tsx for React components

## Architecture Notes
- Backend follows constructor-based dependency injection pattern
- UI uses component-based architecture with centralized API client in `lib/api.ts`
- All UI components should have well-defined prop interfaces