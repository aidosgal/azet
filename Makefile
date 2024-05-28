# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=azet
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_PACKAGE=./cmd/azet

# Default target executed when you run `make`
all: run 

# Build the application
build: 
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Run the application
run: build
	./$(BINARY_NAME)

# Install dependencies
deps:
	$(GOGET) -u ./...

# Test the application
test: 
	$(GOTEST) -v ./...

# Clean build artifacts
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Cross-compile for Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) $(MAIN_PACKAGE)

# Include a help target to describe make targets
help:
	@echo "Makefile commands:"
	@echo "  all          - Runs tests and builds the application"
	@echo "  run          - Runs the application"
	@echo "  build        - Builds the application"
	@echo "  deps         - Installs dependencies"
	@echo "  test         - Runs tests"
	@echo "  clean        - Cleans build artifacts"
	@echo "  build-linux  - Cross-compiles the application for Linux"
