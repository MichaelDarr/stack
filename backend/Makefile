#
# /home/michael/projects/kitchen_sink/web_project/backend
#

APP := backend
BIN := ./bin
GO ?= go

GOFLAGS :=
EXTRA_GOFLAGS ?=

MIGRATION_NAME ?= new_migration
MIGRATION_DIR := $(CURDIR)/db/migrations
POSTGRESQL_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(BACKEND_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)

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
