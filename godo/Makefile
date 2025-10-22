.PHONY: help build run test clean install fmt vet lint all

# Binary name
BINARY_NAME=godo
OUTPUT_DIR=bin

# Version info (can be overridden)
VERSION?=dev
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# Default target
all: fmt vet test build

## help: Display this help message
help:
	@echo "Available targets:"
	@echo "  make build      - Build the binary"
	@echo "  make run        - Run the application"
	@echo "  make test       - Run tests"
	@echo "  make clean      - Remove built binaries"
	@echo "  make install    - Install binary to GOPATH/bin"
	@echo "  make fmt        - Format code"
	@echo "  make vet        - Run go vet"
	@echo "  make lint       - Run golangci-lint (if installed)"
	@echo "  make all        - Format, vet, test, and build"
	@echo ""
	@echo "Build with version:"
	@echo "  make build VERSION=1.0.0"

## build: Build the application binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(OUTPUT_DIR)
	@go build $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME) ./cmd/todo
	@echo "Binary created at $(OUTPUT_DIR)/$(BINARY_NAME)"

## run: Run the application
run:
	@go run ./cmd/todo/main.go $(ARGS)

## test: Run all tests
test:
	@echo "Running tests..."
	@go test -v ./...

## test-coverage: Run tests with coverage report
test-coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## clean: Remove build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(OUTPUT_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

## install: Install the binary to GOPATH/bin
install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(OUTPUT_DIR)/$(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)
	@echo "Installed to $(GOPATH)/bin/$(BINARY_NAME)"

## fmt: Format all Go files
fmt:
	@echo "Formatting code..."
	@go fmt ./...

## vet: Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

## lint: Run golangci-lint (requires golangci-lint to be installed)
lint:
	@if command -v golangci-lint > /dev/null; then \
		echo "Running golangci-lint..."; \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

## deps: Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

## dev: Run in development mode with auto-reload (requires entr)
dev:
	@if command -v entr > /dev/null; then \
		find . -name "*.go" | entr -r go run ./cmd/todo/main.go; \
	else \
		echo "entr not installed. Install it for auto-reload support"; \
		echo "On Ubuntu/Debian: sudo apt-get install entr"; \
	fi
