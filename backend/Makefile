.PHONY: build run test dev db-up db-down db-reset db-shell db-logs

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

# Database commands
db-up:
	@echo "Starting PostgreSQL container..."
	docker-compose -f ../docker-compose.yml up -d postgres

db-down:
	@echo "Stopping PostgreSQL container..."
	docker-compose -f ../docker-compose.yml stop postgres

db-reset:
	@echo "Resetting database..."
	docker-compose -f ../docker-compose.yml down -v postgres
	docker-compose -f ../docker-compose.yml up -d postgres

db-shell:
	@echo "Connecting to database shell..."
	docker-compose -f ../docker-compose.yml exec postgres psql -U zboxes_user -d zboxes_db

db-logs:
	@echo "Showing database logs..."
	docker-compose -f ../docker-compose.yml logs -f postgres
