.PHONY: all help docker-up docker-down docker-build clean install build run dev test

# Default target
all: help

# Install all dependencies
install:
	@echo "Installing all dependencies..."
	@cd backend && go mod download
	@cd frontend && $(MAKE) install

# Build all applications
build:
	@echo "Building all applications..."
	@cd backend && $(MAKE) build
	@cd frontend && $(MAKE) build

# Run all applications in production mode
run:
	@echo "Running all applications..."
	@cd backend && $(MAKE) run &
	@cd frontend && $(MAKE) run

# Run all applications in development mode
dev:
	@echo "Starting development servers..."
	@cd backend && $(MAKE) db-up
	@cd backend && $(MAKE) dev &
	@cd frontend && $(MAKE) dev

# Run all tests
test:
	@echo "Running all tests..."
	@cd backend && $(MAKE) test
	@cd frontend && $(MAKE) test

# Docker commands
docker-up:
	@echo "Starting Docker containers..."
	docker-compose up --build -d

docker-down:
	@echo "Stopping Docker containers..."
	docker-compose down

docker-build:
	@echo "Building Docker images..."
	docker-compose build

# Clean everything
clean:
	@echo "Cleaning up..."
	@cd backend && rm -rf bin/
	@cd frontend && rm -rf dist/ node_modules/
	docker-compose down -v
	docker system prune -f

# Show help
help:
	@echo "Available targets:"
	@echo "  install      - Install all dependencies"
	@echo "  build        - Build all applications"
	@echo "  run          - Run all applications in production mode"
	@echo "  dev          - Run all applications in development mode"
	@echo "  test         - Run all tests"
	@echo "  docker-up    - Start all Docker containers"
	@echo "  docker-down  - Stop all Docker containers"
	@echo "  docker-build - Build Docker images"
	@echo "  clean        - Clean up build artifacts and Docker resources"
	@echo "  help         - Show this help message"
	@echo ""
	@echo "Database commands (run from backend directory):"
	@echo "  backend/make db-up     - Start the database"
	@echo "  backend/make db-down   - Stop the database"
	@echo "  backend/make db-reset  - Reset the database"
	@echo "  backend/make db-shell  - Open database shell"
	@echo "  backend/make db-logs   - Show database logs" 