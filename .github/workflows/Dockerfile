# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Set necessary environment variables
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -v -o tarmac-fox .

# Stage 2: Create the final, minimal image
FROM alpine:latest

# Install ca-certificates for making HTTPS requests from the app
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/tarmac-fox .

# Command to run the application
CMD ["./tarmac-fox"]