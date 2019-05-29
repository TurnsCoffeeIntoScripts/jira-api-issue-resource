PKGS=$(shell go list ./... | grep -v /vendor)

.PHONY: all build test clean run remake

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
JIRA_API_PATH=cmd/jira-api/
JIRA_API_BINARY=jiraApiResource

all: build test

build:
	cd $(JIRA_API_PATH) && $(GOBUILD) -o $(JIRA_API_BINARY) -v

test: build
	$(GOTEST) $(PKGS)

clean:
	cd $(JIRA_API_PATH) && $(GOCLEAN)
	rm $(JIRA_API_PATH)/$(JIRA_API_BINARY)

remake: clean all

run:
	$(JIRA_API_PATH)/$(JIRA_API_BINARY) --username=TestUser1
