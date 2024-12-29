# Use the official Golang image
FROM golang:1.23.4

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Expose application port
EXPOSE 8080

# Start the application
CMD ["./main"]
