package configuration

import (
	"flag"
)

type JiraApiResourceFlags struct {
	ShowHelp   *bool
	JiraApiUrl *string
	Protocol   *string
	Username   *string
	Password   *string
	IssueId    *string
	Body       *string
}

func (f *JiraApiResourceFlags) SetupFlags(parse bool) {
	f.ShowHelp = flag.Bool("help", false, "")
	f.JiraApiUrl = flag.String("url", "", "The base URL of the Jira Rest API to be used (without the http|https)")
	f.Protocol = flag.String("protocol", "https", "The http protocol to be used (http|https)")
	f.Username = flag.String("username", "", "Username used to establish a secure connection with the Jira Rest API")
	f.Password = flag.String("password", "", "Password used by the username in the connection to the Jira Rest API")
	f.IssueId = flag.String("id", "", "The Jira ticket ID (Format: <PROJECT_KEY>-<NUMBER>")
	f.Body = flag.String("body", "", "The body of content to set (description, comment, etc.")

	if parse {
		flag.Parse()
	}
}

func (f *JiraApiResourceFlags) ValidateBaseFlags() bool {
	if *f.IssueId == "" || *f.JiraApiUrl == "" || *f.Username == "" || *f.Password == "" {
		flag.Usage()
		return false
	}

	return true
}