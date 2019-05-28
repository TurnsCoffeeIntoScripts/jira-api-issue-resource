# TODO see: https://sohlich.github.io/post/go_makefile/

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
FIND_TAG_BINARY_PATH=cmd/find-tag/
FIND_TAG_BINARY=findTag

# TODO add test
all: build

build:
	make build_find_tag

build_find_tag:
	cd $(FIND_TAG_BINARY_PATH) && $(GOBUILD) -o $(FIND_TAG_BINARY) -v

#test:

clean:
	cd $(FIND_TAG_BINARY_PATH) && $(GOCLEAN)
	rm $(FIND_TAG_BINARY_PATH)/$(FIND_TAG_BINARY)
