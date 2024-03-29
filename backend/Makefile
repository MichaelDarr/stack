#
# github.com/MichaelDarr/stack/backend
#

APP := backend
MIGRATE_APP := migrate
BIN := ./bin
DESTDIR :=
PREFIX := /usr/local
GO ?= go

GOFLAGS :=
EXTRA_GOFLAGS ?=

GENERATED_RESOLVERS_DIR := $(CURDIR)/internal/resolvers/generated
GENERATED_GQL_DIR := $(CURDIR)/internal/graphql/generated

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
	make build-migrate
	make install-migrate
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

.PHONY: build-migrate
build-migrate:
	$(GO) build $(GOFLAGS) -ldflags '-s -w' $(EXTRA_GOFLAGS) -o $(BIN)/$(MIGRATE_APP) ./cmd/$(MIGRATE_APP)

.PHONY: install-migrate
install-migrate:
	install -d $(DESTDIR)$(PREFIX)/bin
	install -m 0755 $(BIN)/$(MIGRATE_APP) $(DESTDIR)$(PREFIX)/bin/$(MIGRATE_APP)

.PHONY: uninstall-migrate
uninstall-migrate:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(BIN)
