.PHONY: test install dev

GO ?= go

test:
	$(GO) test ./...

install:
	$(GO) install ./...

dev:
	air

build:
	$(GO) build -o ./tmp/api ./cmd/api

build-prod:
	$(GO) build -ldflags "-s -w" -o ./tmp/api-prod ./cmd/api