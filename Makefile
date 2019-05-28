# TODO see: https://sohlich.github.io/post/go_makefile/

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
JIRA_API_PATH=cmd/jira-api/
JIRA_API_BINARY=jiraApiResource

# TODO add test
all: build

build:
	cd $(JIRA_API_PATH) && $(GOBUILD) -o $(JIRA_API_BINARY) -v

#test:

clean:
	cd $(JIRA_API_PATH) && $(GOCLEAN)
	rm $(JIRA_API_PATH)/$(JIRA_API_BINARY)
