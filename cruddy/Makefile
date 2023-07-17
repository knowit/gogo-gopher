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

generate:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@${SQLC_VERSION}
	sqlc generate --file internal/database/sqlc/sqlc.yaml

build:
	mkdir -p bin/
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-${OS} cmd/main.go

run:
	./bin/${BINARY_NAME}-${OS}

clean:
	go clean
	rm -f bin/*

clean-run: clean generate build run
