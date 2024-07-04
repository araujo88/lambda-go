GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

# Default command when running 'make'
all: test build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Build the project for Unix systems
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

# Run tests across all packages
test:
	$(GOTEST) -v -race -cover ./...

# Clean up binaries
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Fetch dependencies
deps:
	$(GOGET) ./...

# Run a specific test pattern
test-pattern:
	$(GOTEST) -v ./... -run $(pattern)
