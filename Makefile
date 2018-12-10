all: install

install:
	go install

fmt:
	go fmt ./...

.PHONY: install fmt
