IMPORT := github.service.anz/ecp/github-enterprise-exporter
UNAME := $(shell uname -s)

# Default to 64bit linux binaries
GOOS ?= linux
GOARCH ?= amd64

ifeq ($(UNAME), Darwin)
	GOBUILD ?= CGO_ENABLED=0 go build -i
else
	GOBUILD ?= CGO_ENABLED=0 go build
endif

PACKAGES ?= $(shell go list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f -not -path "./node_modules/*")

TAGS ?= netgo

ifndef DATE
	DATE := $(shell date -u '+%Y%m%d')
endif

ifndef SHA
	SHA := $(shell git rev-parse --short HEAD)
endif

ifndef VERSION
	ifneq ($(RELEASE_TAG),)
		VERSION ?= $(subst v,,$(RELEASE_TAG))
	else
		VERSION ?= $(SHA)
	endif
endif

# Embed the version and git commit into the built binary
LDFLAGS += -s -w -extldflags "-static" -X "$(IMPORT)/pkg/version.String=$(VERSION)" -X "$(IMPORT)/pkg/version.Revision=$(SHA)" -X "$(IMPORT)/pkg/version.Date=$(DATE)"
GCFLAGS += all=-N -l

.PHONY: all
all: build

.PHONY: sync
sync:
	go mod download

.PHONY: clean
clean:
	go clean -i ./...
	rm -rf bin dist

.PHONY: build
build: bin/ghe-exporter

bin/ghe-exporter: $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./cmd/ghe-exporter
