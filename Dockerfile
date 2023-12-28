# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the local code to the container
COPY . .

# Build the Golang application
RUN go build -o bot_frases .

# Expose the port the app runs on
EXPOSE 7000

# Command to run the executable
CMD ["./bot_frases"]
