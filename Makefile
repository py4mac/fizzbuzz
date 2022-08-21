.PHONY: 

# -----------------------------------------------------------------------------
# Linters https://golangci-lint.run/usage/install/

lint:
	golangci-lint run

# -----------------------------------------------------------------------------
# Test

test: lint
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -func coverage.txt | grep total

# -----------------------------------------------------------------------------
# Doc go install golang.org/x/tools/cmd/godoc@latest

doc:
	$(info http://localhost:6060/pkg/github.com/py4mac/fizzbuzz?m=all)
	godoc -http=:6060

# -----------------------------------------------------------------------------
# Swagger https://github.com/swaggo/swag

swagger:
	$(info Starting swagger generating)
	swag init --dir=cmd,internal/fizzbuzz/domain,internal/fizzbuzz/delivery/http/v1 cmd/main.go

# -----------------------------------------------------------------------------
# Run

run:
	$(info Running application)
	go run ./cmd/main.go --config=./config.local.yaml


# -----------------------------------------------------------------------------
# Docker compose commands

start_local:
	$(info Starting docker local environment)
	docker-compose -f docker-compose.local.yml up -d

stop_local:
	$(info Stopping docker local environment)
	docker-compose -f docker-compose.local.yml stop

start_prod:
	$(info Starting docker production environment)
	docker-compose -f docker-compose.yml up --build -d

stop_prod:
	$(info Stopping docker production environment)
	docker-compose -f docker-compose.yml stop

# -----------------------------------------------------------------------------
# Go migrate https://github.com/golang-migrate/migrate

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