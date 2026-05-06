SHELL := /bin/sh

VERSION ?= dev
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -s -X github.com/ayn2op/discordo/internal/version.Version=$(VERSION) -X github.com/ayn2op/discordo/internal/version.Commit=$(COMMIT) -X github.com/ayn2op/discordo/internal/version.Date=$(DATE)

.PHONY: fmt fmt-check test build lint vulncheck check

fmt:
	gofmt -w .

fmt-check:
	test -z "$$(gofmt -l .)"

test:
	go test ./...

build:
	go build -trimpath -ldflags="$(LDFLAGS)" .

lint:
	golangci-lint run

vulncheck:
	govulncheck ./...

check: fmt-check test build
