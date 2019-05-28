# TODO see: https://sohlich.github.io/post/go_makefile/

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
MAIN_PATH=cmd/jira-api/
BINARY=jiraApiResource

# TODO add test
all: build

build:
	cd $(MAIN_PATH) && $(GOBUILD) -o $(BINARY) -v

#test:

clean:
	cd $(MAIN_PATH) && $(GOCLEAN)
	rm $(MAIN_PATH)/$(BINARY)
