SHELL := /bin/sh

VERSION ?= dev
DIST ?= mtg-fork
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -s -X github.com/ayn2op/discordo/internal/version.Version=$(VERSION) -X github.com/ayn2op/discordo/internal/version.Commit=$(COMMIT) -X github.com/ayn2op/discordo/internal/version.Date=$(DATE) -X github.com/ayn2op/discordo/internal/version.Distribution=$(DIST)
SMOKE_VERSION := smoke-mtg
SMOKE_COMMIT := smoke
SMOKE_DATE := 2000-01-01T00:00:00Z
SMOKE_LDFLAGS := -s -X github.com/ayn2op/discordo/internal/version.Version=$(SMOKE_VERSION) -X github.com/ayn2op/discordo/internal/version.Commit=$(SMOKE_COMMIT) -X github.com/ayn2op/discordo/internal/version.Date=$(SMOKE_DATE) -X github.com/ayn2op/discordo/internal/version.Distribution=$(DIST)

.PHONY: fmt fmt-check test build smoke lint vulncheck check

fmt:
	gofmt -w .

fmt-check:
	test -z "$$(gofmt -l .)"

test:
	go test ./...

build:
	go build -trimpath -ldflags="$(LDFLAGS)" .

smoke:
	tmp="$$(mktemp -d)"; \
	trap 'rm -rf "$$tmp"' EXIT; \
	go build -trimpath -ldflags="$(SMOKE_LDFLAGS)" -o "$$tmp/discordo" .; \
	test "$$("$$tmp/discordo" -version)" = "discordo distribution=$(DIST) version=$(SMOKE_VERSION) commit=$(SMOKE_COMMIT) date=$(SMOKE_DATE)"; \
	go test ./internal/config -run TestLoad

lint:
	golangci-lint run

vulncheck:
	govulncheck ./...

check: fmt-check test build smoke
