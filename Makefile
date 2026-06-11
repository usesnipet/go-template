.PHONY: test install dev build build-prod db-generate

GO ?= go
ATLAS ?= atlas
ATLAS_ENV ?= local

# make db-generate add_users
# make db-generate "add users"
# make db-generate MIGRATION_NAME=add_users
MIGRATION_NAME ?= $(strip $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS)))
RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(RUN_ARGS):
	@:

install:
	$(GO) install ./...

dev:
	air

build:
	$(GO) build -o ./tmp/api ./cmd/api

build-prod:
	$(GO) build -ldflags "-s -w" -o ./tmp/api-prod ./cmd/api

db-generate:
	@set -a && [ -f .env ] && . ./.env; set +a; \
	if [ -z "$(MIGRATION_NAME)" ]; then \
		echo 'usage: make db-generate <name>'; \
		echo '       make db-generate MIGRATION_NAME=<name>'; \
		exit 1; \
	fi; \
	$(ATLAS) migrate diff "$(MIGRATION_NAME)" --env $(ATLAS_ENV)
