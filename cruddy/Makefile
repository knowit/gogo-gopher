BINARY_NAME=cruddy

# CLI versions
SQLC_VERSION=v1.19.1

ifeq '$(findstring ;,$(PATH))' ';'
    detected_OS := Windows
else
    detected_OS := $(shell uname 2>/dev/null || echo Unknown)
    detected_OS := $(patsubst CYGWIN%,Cygwin,$(detected_OS))
    detected_OS := $(patsubst MSYS%,MSYS,$(detected_OS))
    detected_OS := $(patsubst MINGW%,MSYS,$(detected_OS))
endif
ifeq ($(detected_OS),Windows)
	OS := windows
else ifeq ($(detected_OS),Darn)
	OS := darwin
else ifeq ($(detected_OS),Linux)
	OS := linux
else ifeq ($(detected_OS),FreeBSD)
	OS := freebsd
else
    echo "Did not detect a support OS for building the application: $(detected_OS)"
endif

.DEFAULT_GOAL := help ## Prints help for targets with comments
.PHONY: help

generate: ### Generating SQL code
	@echo "Generating SQL code..."
	go install github.com/kyleconroy/sqlc/cmd/sqlc@${SQLC_VERSION}
	sqlc generate --file internal/database/sqlc/sqlc.yaml
	@echo "Finished generating"

build: ### Build application
	@echo "Building application..."
	mkdir -p bin/
	GOARCH=amd64 GOOS=${BINARY_NAME} go build -o bin/${BINARY_NAME}-${OS} cmd/main.go
	@echo "Finished building"

run: ### Run application
	@echo "Running application..."
	./bin/${BINARY_NAME}-${OS}
	@echo "Closing down application. Bye!"

test: ### Run tests
	@echo "Running tests..."
	go test ./...
	@echo "Finished running tests"

vet: ### Vet code
	@echo "Vetting code..."
	go vet ./...
	@echo "Finished vetting code!"

fmt: ### Format code
	@echo "Formatting code..."
	go fmt ./...
	@echo "Finished formatting code"

clean: ### Cleaning binaries
	@echo "Cleaning..."
	go clean
	rm -f bin/*
	@echo "Finished cleaning binaries"

clean-run: clean generate build run ### clean+generate+build+run

# Only works on windows!!!
help: ### Print help
	@grep -E '^[a-zA-Z0-9_-]+:.*?### .*$$' $(MAKEFILE_LIST) \
		| sed -n 's/^\(.*\): \(.*\)###\(.*\)/\1|\3/p' \
		| column -t -s '|'
