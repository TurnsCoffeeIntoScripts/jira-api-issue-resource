package main

import (
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/action"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"os"
)

func main() {
	flags := configuration.JiraApiResourceFlags{}
	flags.SetupFlags(true)
	if !action.IssueExists(flags) {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", "Jira Issue doesn't exist. An error might have occured during the API call")
		os.Exit(1)
	}

	action.CommentOnIssue(flags)
}
