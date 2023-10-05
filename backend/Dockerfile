# Build stage
FROM golang:1.21.1-alpine AS builder

# Set working directory
WORKDIR /app

# Set environment variables
ENV GIN_MODE=release

# Copy all files to the container
COPY . .

# Download dependencies and build the Go application
RUN go mod download
RUN go build -o chatty-backend-go .

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy only the compiled binary from the build stage
COPY --from=builder /app/chatty-backend-go /app/chatty-backend-go
COPY --from=builder /app/.env /app/.env

# Expose the application port
EXPOSE 8080

# Run the application
CMD [ "./chatty-backend-go" ]
