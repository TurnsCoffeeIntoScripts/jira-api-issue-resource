package configuration

import (
	"flag"
)

type JiraApiResourceFlags struct {
	ShowHelp      *bool
	JiraApiUrl    *string
	Protocol      *string
	Username      *string
	Password      *string
	IssueId       *string
	RawIssueList  *string
	Body          *string
	ForceOnParent *bool
	ForceFinish   *bool

	// Context flags
	CtxComment *bool

	// Initialized flags (can't be set via any input flags)
	SingleIssue   bool
	MultipleIssue bool
}

func (f *JiraApiResourceFlags) SetupFlags(parse bool) bool {
	f.ShowHelp = flag.Bool("help", false, "")
	f.JiraApiUrl = flag.String("url", "", "The base URL of the Jira Rest API to be used (without the http|https)")
	f.Protocol = flag.String("protocol", "https", "The http protocol to be used (http|https)")
	f.Username = flag.String("username", "", "Username used to establish a secure connection with the Jira Rest API")
	f.Password = flag.String("password", "", "Password used by the username in the connection to the Jira Rest API")
	f.IssueId = flag.String("id", "", "The Jira ticket ID (Format: <PROJECT_KEY>-<NUMBER>")
	f.RawIssueList = flag.String("ids", "", "")
	f.Body = flag.String("body", "", "The body of content to set (description, comment, etc.")
	f.ForceOnParent = flag.Bool("force-on-parent", false, "")
	f.ForceFinish = flag.Bool("force-finish", false, "Force jira-api-resource to execute every API call before exiting, even if a previous one failed")

	// Context flags
	f.CtxComment = flag.Bool("comment", false, "")

	if parse {
		flag.Parse()
		return f.ValidateBaseFlags()
	}

	return true
}

func (f *JiraApiResourceFlags) ValidateBaseFlags() bool {
	if *f.JiraApiUrl == "" || *f.Username == "" || *f.Password == "" {
		return false
	}

	if *f.RawIssueList != "" {
		f.MultipleIssue = true
	} else if *f.IssueId != "" {
		f.SingleIssue = true
	} else {
		return false
	}

	return true
}
