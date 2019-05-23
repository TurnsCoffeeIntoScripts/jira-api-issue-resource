package main

import (
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/action"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/domain"
	"os"
)

func main() {
	flags := configuration.JiraApiResourceFlags{}
	flags.SetupFlags(true)
	var exists bool
	var issue *domain.Issue
	if exists, issue = action.GetIssue(flags); !exists {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", "Jira Issue doesn't exist. An error might have occured during the API call")
		os.Exit(1)
	}

	if action.HasParent(*issue) {
		flags.IssueId = &issue.Fields.Parent.Key
		if exists, issue = action.GetIssue(flags); !exists {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", "Jira Parent Issue doesn't exist. An error might have occured during the API call")
			os.Exit(1)
		}
	}

	action.CommentOnIssue(flags)
}
