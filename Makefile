GO_BIN ?= go
PACKR_BIN ?= packr

all: install

install:
	$(PACKR_BIN) install

fmt:
	$(GO_BIN) fmt ./...

.PHONY: install fmt
