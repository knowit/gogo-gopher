BINARY_NAME=cruddy

# CLI versions
SQLC_VERSION=v1.19.1

generate:
		go install github.com/kyleconroy/sqlc/cmd/sqlc@${SQLC_VERSION}
		sqlc generate --file internal/database/sqlc/sqlc.yaml

build:
		mkdir -p bin/
		GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin cmd/main.go
		#GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/main.go
		#GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows cmd/main.go

# TODO: Make this OS agnostic
run:
		./bin/${BINARY_NAME}-darwin

clean:
		go clean
		rm -f bin/${BINARY_NAME}-darwin
		rm -f bin/${BINARY_NAME}-linux
		rm -f bin/${BINARY_NAME}-windows

clean-run: clean generate build run
