package action

import (
	"flag"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/domain"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
)

func HasParent(i domain.Issue) bool {
	return i.Fields.Parent != nil
}

func GetIssue(flags configuration.JiraApiResourceFlags) (bool, *domain.Issue) {
	success, md := setup(flags)
	if !success {
		return success, nil
	}

	return rest.Get("issue/??", *flags.IssueId, *md)
}

func CommentOnIssue(flags configuration.JiraApiResourceFlags) bool {
	success, md := setup(flags)
	if !success || *flags.Body == "" {
		flag.Usage()
		return success
	}

	//flags.PopulateMap()
	//return rest.Post("issue/??/comment", *flags.IssueId, *md, flags.Data)
	return false
}

func setup(flags configuration.JiraApiResourceFlags) (bool, *rest.Metadata) {
	//if !flags.Validate() {
		//return false, nil
	//}

	md := rest.Metadata{}
	md.Initialize(flags)

	return true, &md
}
