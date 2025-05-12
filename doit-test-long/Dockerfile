FROM golang:1.21-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Install git (if using modules from private repos)
RUN apk add --no-cache git

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o parking-lot

# ---------- Stage 2: Run ----------
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy built binary from builder
COPY --from=builder /app/parking-lot .

# Expose port (change if needed)
EXPOSE 8080

# Run the binary
CMD ["./parking-lot"]