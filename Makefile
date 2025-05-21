# Name of binary targets
BINARIES := importova exportova

# Default target
.PHONY: all
all: build

# Build all binaries
build:
	for b in $(BINARIES); do \
		go build -o bin/$$b ./cmd/$$b; \
	done

# Run all tests
.PHONY: test
test:
	go test ./...

# Run just unit tests
.PHONY: unit
unit:
	go test ./internal/...

# Lint
.PHONY: lint
lint:
	golangci-lint run ./...

# Format code
.PHONY: fmt
fmt:
	go fmt ./...
	gofmt -w -s .
	go mod tidy
	golangci-lint run --fix ./...

# Clean build artifacts
clean:
	rm -rf bin/*
