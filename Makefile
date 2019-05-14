#
# This make file is based on: https://github.com/thockin/go-build-template/blob/master/Makefile
#

# The binary to build
BIN := jira-api-resource

# Where to push the docker image
REGISTRY ?= thescripter777

# Version strings is based on git tags
VERSION := $(shell git describe --tags --always --dirty)

# Directories containing go code
SRC_DIRS := cmd pkg

ALL_PLATFORMS := linux/amd64 linux/arm linux/arm64 linux/ppc64le linux/s390x

# User should pass GOOD and/or GOARCH
OS := $(if $(GOOS),$(GOOS), $(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH), $(shell go env GOARCH))

BASEIMAGE ?= gcr.io/distroless/static

IMAGE := $(REGISTRY)/$(BIN)
TAG := $(VERSION)__$(OS)_$(ARCH)

BUILD_IMAGE ?= golang:1.12-alpine

all: build

build-%:
	@$(MAKE) build 							\
		--no-print-directory				\
		GOOS=$(firstword $(subst _, ,$*)) 	\
		GOARCH=$(lastword $(subst _, ,$*))

## WORK IN PROGRES... MORE TO COME