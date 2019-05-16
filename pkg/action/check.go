package action

import (
	"flag"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
	"net/http"
)

func Check() bool {
	jiraApiUrl := flag.String("url", "", "The base URL of the Jira Rest API to be used (without the http|https)")
	protocol := flag.String("protocol", "https", "The http protocol to be used (http|https)")
	username := flag.String("username", "", "Username used to establish a secure connection with the Jira Rest API")
	password := flag.String("password", "", "Password used by the username in the connection to the Jira Rest API")
	issueId := flag.String("id", "", "")

	flag.Parse()

	if *issueId == "" || *jiraApiUrl == "" || *username == "" || *password == "" {
		flag.Usage()
		return false
	}

	md := rest.Metadata{}
	md.Url = *jiraApiUrl
	md.Protocol = *protocol
	md.HttpMethod = http.MethodGet
	md.Username = *username
	md.Password = *password

	return rest.IssueExists(md, *issueId)
}
