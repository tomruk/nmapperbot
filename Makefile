GO = go
PREFIX ?= /usr/local/bin

all: build

.PHONY: build
build:
	$(GO) build

.PHONY: install
install: nmapperbot
	mv nmapperbot $(PREFIX)

.PHONY: clean
clean:
	rm -f nmapperbot
