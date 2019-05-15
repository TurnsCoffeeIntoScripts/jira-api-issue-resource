package action

import (
	"errors"
	"flag"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
)

func Check() error {
	jiraApiUrl := flag.String("url", "", "")
	username := flag.String("username", "", "")
	password := flag.String("password", "", "")
	issuePrefix := flag.String("prefix", "", "")
	issueId := flag.String("id", "", "")

	flag.Parse()

	if *issuePrefix == "" || *issueId == "" || *jiraApiUrl == "" || *username == "" || *password == "" {
		err := errors.New("all arguments must be set")
		flag.Usage()
		return err
	}

	md := rest.Metadata{}
	md.Url = *jiraApiUrl
	md.Username = *username
	md.Password = *password

	return rest.GetIssue(md, *issuePrefix, *issueId)
}
