package configuration

import (
	"flag"
	"fmt"
)

type JiraApiResourceConfiguration struct {
	Initialized bool
	Parsed      bool
	Valid       bool
	Flags       JiraApiResourceFlags
	Parameters  JiraApiResourceParameters
}

type JiraApiResourceParameters struct {
	JiraApiUrl  *string
	Protocol    *string
	Username    *string
	Password    *string
	Body        *string
	Label       *string
	IssueId     *string
	IssueList   *string
	IssueScript *string
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

func (conf *JiraApiResourceConfiguration) SetupFlags() bool {
	if !conf.Initialized {
		// Setup context flags (JiraApiResourceContextFlags)
		conf.Flags.ContextFlags.ShowHelp = flag.Bool("help", false, "")
		conf.Flags.ContextFlags.CtxComment = flag.Bool("comment", false, "")
		conf.Flags.ContextFlags.CtxAddLabel = flag.Bool("add-label", false, "")

		// Setup aplication flags (JiraApiResourceApplicationFlags)
		conf.Flags.ApplicationFlags.ForceOnParent = flag.Bool("force-on-parent", false, "")
		conf.Flags.ApplicationFlags.ForceFinish = flag.Bool("force-finish", false, "Force jira-api-resource to execute every API call before exiting, even if a previous one failed")
		// Application flags (JiraApiResourceApplicationFlags) that'll be initialized later
		conf.Flags.ApplicationFlags.SingleIssue = false
		conf.Flags.ApplicationFlags.MultipleIssue = false
		conf.Flags.ApplicationFlags.ZeroIssue = false

		// Setup parameters (JiraApiResourceParameters)
		conf.Parameters.JiraApiUrl = flag.String("url", "", "The base URL of the Jira Rest API to be used (without the http|https)")
		conf.Parameters.Protocol = flag.String("protocol", "https", "The http protocol to be used (http|https)")
		conf.Parameters.Username = flag.String("username", "", "Username used to establish a secure connection with the Jira Rest API")
		conf.Parameters.Password = flag.String("password", "", "Password used by the username in the connection to the Jira Rest API")
		conf.Parameters.Body = flag.String("body", "", "The body of content to set (description, comment, etc.")
		conf.Parameters.Label = flag.String("label", "", "")
		conf.Parameters.IssueId = flag.String("issue-id", "", "")
		conf.Parameters.IssueList = flag.String("issue-list", "", "")
		conf.Parameters.IssueScript = flag.String("issue-script", "", "")

		conf.Initialized = true
	}

	// Parse the flags according to the input parameters
	flag.Parse()

	// Validations
	success, errList := conf.ValidateBaseParameters()

	for err := range errList {
		fmt.Println(err)
	}

	return success
}
