# Start from a minimal base image
FROM golang:1.23 as builder
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Ensure static linking and correct architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./apis/services/akhirah

# Use a lightweight base image for production
FROM alpine:latest
WORKDIR /root/

# Install necessary dependencies (if required)
RUN apk --no-cache add ca-certificates

# Copy the built binary from builder
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x main

# Run the application
CMD ["./main"]