# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOLINT=golangci-lint
GOVET=$(GOCMD) vet

# Binary name
BINARY_NAME=memorycache
BINARY_PATH=./bin/$(BINARY_NAME)

# Main package path
MAIN_PATH=./cmd/memorycache

# Coverage
COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html

# Build flags
LDFLAGS=-ldflags "-s -w"

.PHONY: all build clean test coverage fmt vet lint run help deps tidy bench test-verbose test-race

# Default target
all: clean fmt vet test build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p bin
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_PATH) -v $(MAIN_PATH)
	@echo "Build complete: $(BINARY_PATH)"

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	$(BINARY_PATH)

# Run without building
run-dev:
	@echo "Running in development mode..."
	$(GOCMD) run $(MAIN_PATH)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf bin/
	@rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
	@echo "Clean complete"

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v -cover ./...

# Run tests with race detector
test-race:
	@echo "Running tests with race detector..."
	$(GOTEST) -race -v ./...

# Run tests with verbose output
test-verbose:
	@echo "Running tests with verbose output..."
	$(GOTEST) -v -count=1 ./...

# Generate test coverage report
coverage:
	@echo "Generating coverage report..."
	$(GOTEST) -v -coverprofile=$(COVERAGE_FILE) ./...
	$(GOCMD) tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report generated: $(COVERAGE_HTML)"

# Show coverage in terminal
coverage-term:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=$(COVERAGE_FILE) ./...
	$(GOCMD) tool cover -func=$(COVERAGE_FILE)

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem ./...

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .
	@echo "Format complete"

# Run go vet
vet:
	@echo "Running go vet..."
	$(GOVET) ./...
	@echo "Vet complete"

# Run linter (requires golangci-lint to be installed)
lint:
	@echo "Running linter..."
	@if command -v $(GOLINT) > /dev/null; then \
		$(GOLINT) run ./...; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOGET) -v ./...
	@echo "Dependencies downloaded"

# Tidy go modules
tidy:
	@echo "Tidying modules..."
	$(GOMOD) tidy
	@echo "Modules tidied"

# Install development tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Tools installed"

# Docker targets (for future containerization)
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME):latest .

docker-run:
	@echo "Running Docker container..."
	docker run --rm -p 8080:8080 $(BINARY_NAME):latest

# Help target
help:
	@echo "Available targets:"
	@echo "  make build         - Build the binary"
	@echo "  make run           - Build and run the application"
	@echo "  make run-dev       - Run without building (using go run)"
	@echo "  make test          - Run tests"
	@echo "  make test-race     - Run tests with race detector"
	@echo "  make test-verbose  - Run tests with verbose output"
	@echo "  make coverage      - Generate test coverage report"
	@echo "  make coverage-term - Show coverage in terminal"
	@echo "  make bench         - Run benchmarks"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make fmt           - Format code"
	@echo "  make vet           - Run go vet"
	@echo "  make lint          - Run linter (requires golangci-lint)"
	@echo "  make deps          - Download dependencies"
	@echo "  make tidy          - Tidy go modules"
	@echo "  make install-tools - Install development tools"
	@echo "  make all           - Clean, format, vet, test, and build"
	@echo "  make help          - Show this help message"
