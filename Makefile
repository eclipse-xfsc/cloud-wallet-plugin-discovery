NAME := plugin-discovery
VERSION := $$(git rev-parse HEAD | cut -c -5)
GOVERSION := $(shell go version)
BUILDDATE := $(shell date -u +"%B %d, %Y")
PKG_RELEASE ?= 1
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.buildDate=$(BUILDDATE)' \
           -X 'main.buildGoVersion=$(GOVERSION)'

# development tasks
PACKAGES := $$(go list ./... | grep -v /vendor/ | grep -v /cmd/)

# build tasks
SOURCES := $(shell find . -type f \( -name '*.go' -and -not -name '*_test.go' \))
build: $(SOURCES)
	go build -ldflags "$(LDFLAGS)" -o build/$(NAME)

run:build
	build/plugin-discovery

clean:
	go clean