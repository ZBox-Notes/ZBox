.PHONY: build run test dev

# Build the Go application
build:
	@echo "Building backend..."
	go build -o bin/server

# Run the application locally
run:
	@echo "Running backend locally..."
	go run server.go

# Run tests
test:
	@echo "Running backend tests..."
	go test ./... -v

# Development with hot reload (requires air: go install github.com/cosmtrek/air@latest)
dev:
	@echo "Starting backend development server with hot reload..."
	air 