# SMTP Sink UI

The UI component for the SMTP Sink application, built with React, TypeScript, and Vite.

## Features

- Display received emails in a user-friendly interface
- View email details including headers, text body, and HTML body
- View and download email attachments
- Real-time updates (polling every 5 seconds)
- Clear all emails

## Development

### Prerequisites

- Node.js 18+ 
- npm or yarn

### Getting Started

1. Install dependencies:

```bash
npm install
```

2. Start the development server:

```bash
npm run dev
```

This will start the UI in development mode with HMR (Hot Module Replacement). The UI will proxy API requests to the Go backend running on port 8080.

### Building

To build the UI for production:

```bash
npm run build
```

This will create a production build in the `dist` directory, which can be embedded in the Go server.

## API Endpoints

The UI connects to the following API endpoints:

- `GET /api/emails` - Get all emails
- `GET /api/emails/:id` - Get a specific email
- `DELETE /api/emails` - Clear all emails
- `GET /api/info` - Get server information

Make sure the Go server is running on port 8080 while developing the UI.