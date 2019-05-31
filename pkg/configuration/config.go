package configuration

import (
	"flag"
)

type JiraApiResourceConfiguration struct {
	Parsed     bool
	Valid      bool
	Flags      JiraApiResourceFlags
	Parameters JiraApiResourceParameters
}

type JiraApiResourceParameters struct {
	JiraApiUrl   *string
	Protocol     *string
	Username     *string
	Password     *string
	Body         *string
	Label        *string
	IssueId      *string
	RawIssueList *string
	IssueScript  *string
}

type JiraApiResourceFlags struct {
	ApplicationFlags JiraApiResourceApplicationFlags
	ContextFlags     JiraApiResourceContextFlags
}

type JiraApiResourceApplicationFlags struct {
	ForceOnParent *bool
	ForceFinish   *bool
	SingleIssue   bool
	MultipleIssue bool
	ZeroIssue     bool
}

type JiraApiResourceContextFlags struct {
	ShowHelp    *bool
	CtxComment  *bool
	CtxAddLabel *bool
}

func (f *JiraApiResourceFlags) SetupFlags(parse bool) bool {
	f.ShowHelp = flag.Bool("help", false, "")
	f.JiraApiUrl = flag.String("url", "", "The base URL of the Jira Rest API to be used (without the http|https)")
	f.Protocol = flag.String("protocol", "https", "The http protocol to be used (http|https)")
	f.Username = flag.String("username", "", "Username used to establish a secure connection with the Jira Rest API")
	f.Password = flag.String("password", "", "Password used by the username in the connection to the Jira Rest API")
	f.Body = flag.String("body", "", "The body of content to set (description, comment, etc.")
	f.Label = flag.String("label", "", "")
	f.ForceOnParent = flag.Bool("force-on-parent", false, "")
	f.ForceFinish = flag.Bool("force-finish", false, "Force jira-api-resource to execute every API call before exiting, even if a previous one failed")

	// Issues
	f.IssueId = flag.String("id", "", "The Jira ticket ID (Format: <PROJECT_KEY>-<NUMBER>")
	f.RawIssueList = flag.String("ids", "", "")
	f.IssueScript = flag.String("script", "", "")

	// Context flags
	f.CtxComment = flag.Bool("comment", false, "")
	f.CtxAddLabel = flag.Bool("add-label", false, "")

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
		f.ZeroIssue = true
	}

	return true
}
