#
# /home/michael/projects/shelf/auth
#

APP := auth
BIN := ./bin
GO ?= go

GOFLAGS :=
EXTRA_GOFLAGS ?=

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

.PHONY: gen
gen:
	cd $(CURDIR)
	protoc --go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/auth.proto
