# TODO see: https://sohlich.github.io/post/go_makefile/

PKGS :=	$(shell go list ./... | grep -v /vendor)

.PHONY: test

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
JIRA_API_PATH=cmd/jira-api/
JIRA_API_BINARY=jiraApiResource
BASE_PACKAGE_NAME=github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg

# TODO add test
all: build test

build:
	cd $(JIRA_API_PATH) && $(GOBUILD) -o $(JIRA_API_BINARY) -v

test:
	go test $(PKGS)

clean:
	cd $(JIRA_API_PATH) && $(GOCLEAN)
	rm $(JIRA_API_PATH)/$(JIRA_API_BINARY)

run:
	$(JIRA_API_PATH)/$(JIRA_API_BINARY) --username=TestUser1
