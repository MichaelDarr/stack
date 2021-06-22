#
# /home/michael/projects/shelf/backend
#

APP := backend
BIN := ./bin
GO ?= go

GOFLAGS :=
EXTRA_GOFLAGS ?=

GENERATED_RESOLVERS_DIR := $(CURDIR)/internal/resolvers/generated
GENERATED_GQL_DIR := $(CURDIR)/internal/graphql/generated

MIGRATION_NAME ?= new_migration
MIGRATION_DIR := $(CURDIR)/db/migrations
POSTGRESQL_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

.PHONY: default
default: $(APP)

.PHONY: build
build: $(APP)

.PHONY: $(APP)
$(APP): ## build
	$(GO) build $(GOFLAGS) -ldflags '-s -w' $(EXTRA_GOFLAGS) -o $(BIN)/$@ ./cmd/$(APP)

.PHONY: dev
dev: ## run in dev mode (autoreloading enabled, requires gin)
	go mod download
	gin --immediate \
		--bin=$(BIN)/$(APP) \
		--build ./cmd/$(APP) \
		--buildArgs="$(GOFLAGS) $(EXTRA_GOFLAGS)" \
		run 

.PHONY: generate
generate:
	rm -rf ${GENERATED_RESOLVERS_DIR}/*
	rm -rf ${GENERATED_GQL_DIR}/*
	$(GO) run -v github.com/99designs/gqlgen

.PHONY: downgrade
downgrade:
	migrate -database $(POSTGRESQL_URL) -path $(MIGRATION_DIR) down $(n)

.PHONY: upgrade
upgrade:
	migrate -database $(POSTGRESQL_URL) -path $(MIGRATION_DIR) up $(n)

.PHONY: create
create:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(MIGRATION_NAME)
