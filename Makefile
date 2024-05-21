APP_NAME = tomodachi-life-server
SRC_DIR = .
BUILD_DIR = ./bin

GO = go
GOFMT = gofmt
GOTEST = $(GO) test
GOBUILD = $(GO) build
GOCLEAN = $(GO) clean
GOVET = $(GO) vet
GOMOD = $(GO) mod

# Flags
BUILD_FLAGS = -o $(BUILD_DIR)/$(APP_NAME)

.PHONY: all fmt vet test build run clean

all: fmt vet test build

fmt:
	@echo "Running gofmt..."
	$(GOFMT) -w $(SRC_DIR)

vet:
	@echo "Running go vet..."
	$(GOVET) ./...

test:
	@echo "Running tests..."
	$(GOTEST) ./...

build:
	@echo "Building the application..."
	$(GOBUILD) $(BUILD_FLAGS) $(SRC_DIR)/main.go

run: build
	@echo "Running the application..."
	$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

mod-tidy:
	@echo "Tidying up go.mod..."
	$(GOMOD) tidy

mod-download:
	@echo "Downloading dependencies..."
	$(GOMOD) download
