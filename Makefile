# Variables
GO := go
GOMOD := $(GO) mod
GOBUILD := $(GO) build
GOTEST := $(GO) test
GOCLEAN := $(GO) clean
GOLINT := golangci-lint

# Run the API
run:
	$(GO) run cmd/server/main.go

# Install dependencies
install:
	$(GO) mod tidy

# Format the code
fmt:
	$(GO) fmt ./...

# Run tests
test:
	$(GOTEST) ./...

# Migrate database (if using a migration tool)
migrate:
	$(info ********************* automigrating the model *************************)
	$(GO) run cmd/migrate/main.go

# Lint the code
lint:
	$(GOLINT) run ./...

# Build the application
build:
	$(GOBUILD) -o ginco-server cmd/server/main.go
	$(GOBUILD) -o ginco-migrate cmd/migrate/main.go

# Clean build files
clean:
	$(GOCLEAN)

lint:
	golangci-lint run ./...

.PHONY: run install fmt test migrate lint build clean
