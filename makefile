.PHONY: default run build test docs clean

# Variables
APP_NAME=api-locadora

APP_VERSION=1.0.0

APP_PORT=8080

# Default target
default: run

# Run the application
run:
	@echo "Running $(APP_NAME) version $(APP_VERSION) on port $(APP_PORT)..."
	@go run ./cmd/main.go