GO ?= go
CC ?= gcc
ifeq ($(GOPATH),)
export GOPATH := $(shell $(GO) env GOPATH)
endif
FIRST_GOPATH := $(firstword $(subst :, ,$(GOPATH)))
GOBIN := $(shell $(GO) env GOBIN)
ifeq ($(GOBIN),)
	GOBIN := $(FIRST_GOPATH)/bin
endif

all: build test magicgen

lint: 
	@golint magic_gen.go

build:
	 @$(GO) build -v

test/magicgen:
	@rm -f testmagic.json 
	@$(GO) run magic_gen.go magic.h testmagic.json

test: test/magicgen
		
magicgen:
	go run magic_gen.go magic.h magic.json

install.tools:
	if [ ! -x "$(GOBIN)/golint" ]; then \
		$(GO) get -u go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0; \
	fi

clean:
	@rm -f testmagic.json

.PHONY: build test clean lint magicgen