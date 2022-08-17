SHELL=/bin/bash -o pipefail

.PHONY: 

# ==============================================================================
# Linters https://golangci-lint.run/usage/install/

lint:
	golangci-lint run

# ==============================================================================
# Test

test: lint
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -func coverage.txt | grep total

# ==============================================================================
# Doc

doc:
	$(info http://localhost:6060/pkg/github.com/py4mac/fizzbuzz)
	godoc -http=:6060

# ==============================================================================
# Swagger https://github.com/swaggo/swag

swagger:
	echo "Starting swagger generating"
	swag init -g **/**/*.go

# ==============================================================================
# Run

run:
	go run ./cmd/main.go --config=./config.yaml


# ==============================================================================
# Docker compose commands

develop:
	echo "Starting docker environment"
	docker-compose -f docker-compose.dev.yml up --build

# ==============================================================================
# Go migrate eduterm-pgql https://github.com/golang-migrate/migrate

DB_NAME = fizzbuzz
DB_HOST = localhost
DB_PORT = 5432
SSL_MODE = disable

force_db:
	migrate -database postgres://postgres:admin@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations force 1

version_db:
	migrate -database postgres://postgres:admin@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations version

migrate_up:
	migrate -database postgres://postgres:admin@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations up 1

migrate_down:
	migrate -database postgres://postgres:admin@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations down 1