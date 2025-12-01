APP_NAME := mini-geth

.PHONY: all build run test clean

all: build run

build:
	@echo "Building..."
	go build -o bin/$(APP_NAME) ./...

run:
	@echo "Running..."
	go run ./...

test:
	@echo "Running tests..."
	go test ./... -v

clean:
	@echo "Cleaning..."
	rm -rf bin/
