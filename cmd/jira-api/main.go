package main

import (
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/action"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/google/uuid"
)

func main() {
	flags := configuration.JiraApiResourceFlags{}
	action.CommentOnIssue(flags)

	fmt.Println(uuid.New())
}
