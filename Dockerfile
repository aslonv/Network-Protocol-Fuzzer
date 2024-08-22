# Use the official Golang image as the base image
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o fuzzer ./cmd/fuzzer

# Use a smaller image for the final stage
FROM alpine:latest  

# Install necessary libraries
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/fuzzer .

# Command to run the executable
ENTRYPOINT ["./fuzzer"]

# Expose the port that the application will run on
EXPOSE 8080