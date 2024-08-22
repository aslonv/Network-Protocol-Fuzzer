# Start with a Go base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o fuzzer ./cmd/fuzzer

# Command to run the executable
CMD ["./fuzzer", "--protocol=protocol.json"]
