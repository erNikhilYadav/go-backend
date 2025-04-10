# Use the official Go image as a builder
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Use a smaller image for the final container
FROM alpine:latest

# Install SQLite and create necessary directories
RUN apk add --no-cache sqlite && \
    mkdir -p /app/data && \
    chmod 777 /app/data

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Create database directory and set permissions
RUN mkdir -p /app/data && \
    chmod 777 /app/data

# Expose the port
EXPOSE 8080

# Set default environment variables
ENV ENVIRONMENT=uat
ENV PORT=8080
ENV DATABASE_DIR=/app/data

# Run the application
CMD ["./main"] 