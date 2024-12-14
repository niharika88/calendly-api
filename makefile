# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOCLEAN := $(GOCMD) clean
GOGET := $(GOCMD) get
BINARY_NAME := calendly-api
PKG_DIR:= ./cmd/.

# Build settings
BUILD_DIR := ./build
BUILD_OUTPUT := $(BUILD_DIR)/$(BINARY_NAME)

build: clean 
	$(GOBUILD) -o $(BUILD_OUTPUT) $(PKG_DIR)

run: build 
	$(BUILD_OUTPUT)

test: 
	$(GOTEST) -v -cover ./...

clean: 
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

gen-swagger:
	swag fmt -d ./cmd,./internal
	swag init -g ./cmd/main.go -pd

help: 
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'
