GO ?= "go"

.PHONY: all
all: lint test

.PHONY: lint
lint:
	@echo "==> $@"
	@$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint run ./...

.PHONY: test
test:
	@echo "==> $@"
	@$(GO) test ./...