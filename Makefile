GO_BIN ?= go

XDG_CONFIG_HOME ?= $(HOME)/.config

NBSPIFY_NAME ?= nbspify

NBSPIFY_CONFIG_DIR ?= $(XDG_CONFIG_HOME)/$(NBSPIFY_NAME)
NBSPIFY_DEFAULTS_DIR ?= ./_defaults
NBSPIFY_DEFINITIONS_FILE_NAME ?= definitions.yaml

all: install

install:
	$(GO_BIN) install

config:
	install -Dm644 $(NBSPIFY_DEFAULTS_DIR)/$(NBSPIFY_DEFINITIONS_FILE_NAME) $(NBSPIFY_CONFIG_DIR)/$(NBSPIFY_DEFINITIONS_FILE_NAME)

fmt:
	$(GO_BIN) fmt ./...

test:
	$(GO_BIN) test -cover -v ./...

test-ci:
	$(GO_BIN) test -v -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: install config fmt test
