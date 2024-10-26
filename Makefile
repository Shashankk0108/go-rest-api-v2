.PHONY: help install build run dev test test-coverage lint clean docker-build docker-up docker-down

help: ## Display this help screen
	@echo "Available commands:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Install dependencies
	go mod download
	go mod tidy
	@echo "Dependencies installed successfully"

build: ## Build the application
	go build -o bin/main .
	@echo "Build completed successfully"

run: build ## Run the application
	./bin/main

dev: ## Run the application in development mode
	go run main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"

lint: ## Run linter
	go fmt ./...
	go vet ./...

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f coverage.out coverage.html
	@echo "Cleaned build artifacts"

docker-build: ## Build Docker image
	docker build -t go-rest-api:latest .
	@echo "Docker image built successfully"

docker-up: ## Start Docker containers
	docker-compose up -d
	@echo "Docker containers started"

docker-down: ## Stop Docker containers
	docker-compose down
	@echo "Docker containers stopped"

docker-logs: ## View Docker logs
	docker-compose logs -f

docker-restart: docker-down docker-up ## Restart Docker containers
