SHELL=/bin/bash -o pipefail

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: fmt lint
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -func coverage.txt | grep total

.PHONY: doc
doc:
	$(info http://localhost:6060/pkg/github.com/py4mac/fizzbuzz)
	godoc -http=:6060

.PHONY: run
run:
	go run main.go serve --port=":8000" --timeout=300