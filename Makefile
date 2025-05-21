# Name of binary targets
BINARIES := importova exportova

# Default target
.PHONY: all
all: build

# Build all binaries
build:
	@for b in $(BINARIES); do \
		go build -o bin/$$b ./cmd/$$b; \
	done
	@echo "Binaries built in bin/"

# Run all tests
.PHONY: test
test:
	go test ./...
	@echo "All tests passed"

# Run just unit tests
.PHONY: unit
unit:
	@go test ./internal/...
	@echo "All unit tests passed"
# Lint
.PHONY: lint
lint:
	@golangci-lint run ./...
	@echo "Linting passed"

# Format code
.PHONY: fmt
fmt:
	@go fmt ./...
	@gofmt -w -s .
	@go mod tidy
	@golangci-lint run --fix ./...
	@echo "Code formatted"

# Clean build artifacts
clean:
	@rm -rf bin/*
	@echo "Cleaned build artifacts"
