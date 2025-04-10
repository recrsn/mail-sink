### Build UI stage ###
FROM node:20-slim AS ui-builder
WORKDIR /app/ui

# Copy UI package files first for better caching
COPY ui/package.json ui/package-lock.json* ./

# Install dependencies
RUN npm ci

# Copy UI source files
COPY ui/ ./

# Build UI
RUN npm run build

### Build Go application stage ###
FROM golang:1.24-alpine AS go-builder
WORKDIR /app

# Copy Go module files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy UI build artifacts
COPY --from=ui-builder /app/ui/dist ./ui/dist

# Copy Go source code
COPY . .

# Build Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o mail-sink .

### Final minimal image ###
FROM alpine:3.19

# Add ca-certificates and timezone info
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN adduser -D -u 1000 appuser

WORKDIR /app

# Copy binary from builder
COPY --from=go-builder /app/mail-sink .

# Use non-root user
USER appuser

# Expose ports
EXPOSE 1025 8080

# Set default environment variables
ENV SMTP_PORT=1025
ENV HTTP_PORT=8080
ENV GIN_MODE=release

# Run the application
CMD ["/app/mail-sink"]
