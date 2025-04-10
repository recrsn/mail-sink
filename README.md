# Mail Sink

A simple SMTP server for testing email flows locally. It provides a REST API to access received emails and includes a web UI.

## Features

- SMTP server that captures all incoming emails
- Parses email content including HTML, plain text, and attachments
- REST API to retrieve emails and clear the mailbox
- Web UI for viewing and managing emails
- Single binary deployment with embedded UI

## Usage

### Building

You can use the provided Makefile to build both the UI and the Go binary in one step:

```bash
# Build everything
make

# Or separately
cd ui
npm install
npm run build
cd ..
go build -o mail-sink
```

### Running

Using the Makefile:

```bash
# Run with default settings
make run

# Run with custom ports
make run-custom SMTP_PORT=2525 HTTP_PORT=3000
```

Or directly:

```bash
# Run with default settings
./mail-sink

# Run with custom ports
./mail-sink --smtp-port 2525 --http-port 3000
```

### Development

To run the UI in development mode with hot reloading:

```bash
# Start the Go server
./mail-sink

# In another terminal, start the UI development server
cd ui
npm run dev
```

This will start the UI development server that proxies API requests to the Go backend.

### Command line options

- `--smtp-port PORT`: SMTP server port (default: 1025)
- `--http-port PORT`: HTTP server port (default: 8080)
- `--help`: Show help

## API

The following API endpoints are available:

- `GET /api/emails`: Get all received emails
- `GET /api/emails/:id`: Get a specific email by ID
- `DELETE /api/emails`: Clear all emails
- `GET /api/info`: Get server information (ports)

## SMTP Configuration

To send emails to the Mail Sink, configure your email client or application with:

- SMTP Server: localhost
- SMTP Port: 1025 (or custom port)
- Authentication: None required (all authentication attempts are accepted)

## License

MIT