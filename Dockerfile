# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files from backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy source code from backend
COPY backend/ .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/bin/server ./bin/server

# Copy migrations
COPY --from=builder /app/migrations ./migrations

# Expose port
EXPOSE 8080

# Run the binary directly (Railway will set PORT automatically)
CMD ["./bin/server"]

