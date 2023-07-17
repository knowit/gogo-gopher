.PHONY: generate

generate:
		go install github.com/kyleconroy/sqlc/cmd/sqlc@1.19.1
		sqlc generate --file internal/database/sqlc/sqlc.yaml
