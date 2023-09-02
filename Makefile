#/bin/bash
PWD ?= $(shell pwd)
export PATH := $(PWD)/bin:$(PATH)

GOOS ?= linux
GOARCH ?= amd64

ifeq ($(shell uname -s),Darwin)
	GOOS = darwin
endif

ifeq ($(shell uname -m),arm64)
	GOARCH = arm64
endif

### function for get build os and arch args
### param: $(1) OS
### param: $(2) ARCH
go_build_args = GOOS=$(1) GOARCH=$(2)

default: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

.PHONY: fmt
fmt: ## Format all go codes
	./scripts/goimports-reviser.sh

build: ## Build binary
	@mkdir -p bin
	@$(call go_build_args,$(GOOS),$(GOARCH)) CGO_ENABLED=0 go build -o bin/im-to-notion github.com/ronething/im-to-notion/cmd/app
.PHONY: build

lint: ## Apply go lint check
	@golangci-lint run --timeout 10m ./...
.PHONY: lint

set-e2e-goos:
	$(eval GOOS=linux)
	@echo "e2e GOOS: $(GOOS)"
.PHONY: set-e2e-goos

build-docker-image: set-e2e-goos build
	@docker build -f build/Dockerfile -t ronething/im-to-notion:dev .
.PHONY: build-docker-image
