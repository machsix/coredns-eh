# Makefile for building CoreDNS
BINARY:=coredns wgsd-client
OUTPUT_DIR?=build
SYSTEM:=
CHECKS:=check
BUILDOPTS?=-v
GOPATH?=$(HOME)/go
MAKEPWD:=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
CGO_ENABLED?=0
GOLANG_VERSION ?= $(shell cat .go-version)

export GOSUMDB = sum.golang.org
export GOTOOLCHAIN = go$(GOLANG_VERSION)

.PHONY: all
all: coredns

.PHONY: coredns
coredns:
	mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=$(CGO_ENABLED) $(SYSTEM) go build $(BUILDOPTS) -ldflags="-s -w" -o $(OUTPUT_DIR) ./...

.PHONY: pb
pb:
	$(MAKE) -C pb

.PHONY: clean
clean:
	go clean
	rm -rf $(OUTPUT_DIR)
