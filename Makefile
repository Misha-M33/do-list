# Load env variales from .env if it exist
ifneq (,$(wildcard .env))
  ENV := $(PWD)/.env
  include $(ENV)
endif

SHELL = /bin/bash

ifeq ($(uname -m),$(filter $( uname -m), Darwin x86_64))
  GOARCH ?= amd64
else
  GOARCH ?= $(shell uname -m)
endif

GO111MODULE := on
CGO_ENABLED ?= 0
GOOS ?= linux
TARGET := dolist
# BUILD_VERSION:=$(GIT_URL)/backend/config.buildVersion=$(COMMIT_TAG):$(COMMIT_SHORT_SHA)
# BUILD_DATE:=$(GIT_URL)/backend/config.buildDate=`date "+%Y.%m.%d_%H:%M:%S"`

.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n\033[36m\033[0m"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: lint
linters: ## Run linter
	golangci-lint version
	golangci-lint run

.PHONY: tidy
tidy: lint ## Run go mod tidy
	go mod tidy

.PHONY: download
download: ## Run go mod download
	go mod download

.PHONY: clean
clean: ## Remove temporary files
	go clean -i all
	@rm -rf ./backend/bin/$(TARGET)

.PHONY: build
build: clean tidy ## Build service
	GOOS=${GOOS} GOARCH=${GOARCH} GO111MODULE=${GO111MODULE} CGO_ENABLED=${CGO_ENABLED} \
		go build -ldflags "-X $(BUILD_VERSION) -X $(BUILD_DATE)" -o bin/$(TARGET) ./cmd/$(TARGET)/

.PHONY: vendors
vendors: ## Run go mod vendor
	go mod vendor

.PHONY: up
up: ## Run docker compose up -d
	docker compose -f docker-compose.local.yml down && docker compose -f docker-compose.local.yml up -d

.PHONY: down
down: ## Run docker compose down
	docker compose -f docker-compose.local.yml down
