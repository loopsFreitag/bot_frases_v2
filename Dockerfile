# Build stage
FROM golang:alpine AS builder

# Install necessary dependencies
RUN apk --no-cache add git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o bot_frases .

# Final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/bot_frases .

# Expose the port the app runs on
EXPOSE 7000

# Command to run the application
CMD ["./bot_frases"]

