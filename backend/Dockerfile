# Use the official Golang image to create a build artifact.
# This is the build stage.
FROM golang:1.19 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o example ./backend/example.go

# Start a new stage from scratch
FROM debian:bullseye-slim

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/example /app/example

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/example"]