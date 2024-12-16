.PHONY: all help docker-up docker-down docker-build db-up db-down clean install build run dev test

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

# Database commands
db-up:
	@echo "Starting PostgreSQL container..."
	docker-compose up -d postgres

db-down:
	@echo "Stopping PostgreSQL container..."
	docker-compose stop postgres

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
	@echo "  db-up        - Start only the PostgreSQL container"
	@echo "  db-down      - Stop only the PostgreSQL container"
	@echo "  clean        - Clean up build artifacts and Docker resources"
	@echo "  help         - Show this help message" 