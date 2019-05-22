package action

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
)

func IssueExists(flags configuration.JiraApiResourceFlags) bool {
	success, md := setup(flags)
	if !success {
		return success
	}

	found, issue := rest.Get("issue/??", *flags.IssueId, *md)

	return found && issue != nil
}

func CommentOnIssue(flags configuration.JiraApiResourceFlags) bool {
	success, md := setup(flags)
	if !success || *flags.RawData == "" {
		return success
	}

	flags.PopulateMap()
	return rest.Post("issue/??/comment", *flags.IssueId, *md, flags.Data)
}

func setup(flags configuration.JiraApiResourceFlags) (bool, *rest.Metadata) {
	if !flags.Validate() {
		return false, nil
	}

	md := rest.Metadata{}
	md.Initialize(flags)

	return true, &md
}
