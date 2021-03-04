BUNDLEDESTDIR := public
BUNDLENAME := bundle.js
CONFIG := tsconfig.json
DENO ?= deno
SRCDIR := src
SRCROOT := $(SRCDIR)/index.tsx
SRVROOT := srv/index.ts

DENOFLAGS := --config $(CONFIG)

.PHONY: default
default: dev

.PHONY: dev
dev:
	make -j serve watch

.PHONY: build
build:
	$(DENO) bundle $(DENOFLAGS) $(SRCROOT) $(BUNDLEDESTDIR)/$(BUNDLENAME) --unstable

.PHONY: cache
cache:
	$(DENO) cache $(DENOFLAGS) $(SRCROOT) --unstable

.PHONY: format
format:
	$(DENO) fmt

.PHONY: lint
lint:
	$(DENO) lint $(SRCDIR) --unstable

.PHONY: serve
serve:
	$(DENO) run --allow-net --allow-read $(SRVROOT)

.PHONY: watch
watch:
	$(DENO) bundle $(DENOFLAGS) $(SRCROOT) $(BUNDLEDESTDIR)/$(BUNDLENAME) --unstable --watch
