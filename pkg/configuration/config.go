package configuration

import (
	"flag"
)

// JiraAPIResourceConfiguration defines a container for both flags (true/false) and parameters.
// The internal 'Initialized' flag is set to true when 'SetupFlags' is called for the first time
// The internal 'Parsed' flag is set to true when 'flag.Parse()' is called for the first time
// The internal 'Valid' flag returns whether or not the configuration is in a valid state
// The internal 'Errors' slice contains all error recorded during the validation process
type JiraAPIResourceConfiguration struct {
	Initialized bool
	Parsed      bool
	Valid       bool
	Errors      []error
	Flags       JiraAPIResourceFlags
	Parameters  JiraAPIResourceParameters
}

type JiraAPIResourceParameters struct {
	JiraAPIUrl  *string
	Protocol    *string
	Username    *string
	Password    *string
	Body        *string
	Label       *string
	IssueID     *string
	IssueList   *string
	IssueScript *string
}

type JiraAPIResourceFlags struct {
	ApplicationFlags JiraAPIResourceApplicationFlags
	ContextFlags     JiraAPIResourceContextFlags
}

type JiraAPIResourceApplicationFlags struct {
	ForceOnParent *bool
	ForceFinish   *bool
	SingleIssue   bool
	MultipleIssue bool
	ZeroIssue     bool
}

type JiraAPIResourceContextFlags struct {
	ShowHelp    JiraAPIResourceContextFlagDefinition
	CtxComment  JiraAPIResourceContextFlagDefinition
	CtxAddLabel JiraAPIResourceContextFlagDefinition
}

type JiraAPIResourceContextFlagDefinition struct {
	Value *bool
	Func  string
}

func (conf *JiraAPIResourceConfiguration) SetupFlags() bool {
	if !conf.Initialized {
		// Setup context flags (JiraAPIResourceContextFlags)
		conf.Flags.ContextFlags.ShowHelp.Value = flag.Bool("help", false, "")
		conf.Flags.ContextFlags.ShowHelp.Func = "ValidateShowHelpContext"
		conf.Flags.ContextFlags.CtxComment.Value = flag.Bool("comment", false, "")
		conf.Flags.ContextFlags.CtxComment.Func = "ValidateCommentContext"
		conf.Flags.ContextFlags.CtxAddLabel.Value = flag.Bool("add-label", false, "")
		conf.Flags.ContextFlags.CtxAddLabel.Func = "ValidateAddLabelContext"

		// Setup aplication flags (JiraAPIResourceApplicationFlags)
		conf.Flags.ApplicationFlags.ForceOnParent = flag.Bool("force-on-parent", false, "")
		conf.Flags.ApplicationFlags.ForceFinish = flag.Bool("force-finish", false, "Force jira-api-resource to execute every API call before exiting, even if a previous one failed")
		// Application flags (JiraAPIResourceApplicationFlags) that'll be initialized later (during context validation)
		conf.Flags.ApplicationFlags.SingleIssue = false
		conf.Flags.ApplicationFlags.MultipleIssue = false
		conf.Flags.ApplicationFlags.ZeroIssue = false

		// Setup parameters (JiraAPIResourceParameters)
		conf.Parameters.JiraAPIUrl = flag.String("url", "", "The base URL of the Jira Rest API to be used (without the http|https)")
		conf.Parameters.Protocol = flag.String("protocol", "https", "The http protocol to be used (http|https)")
		conf.Parameters.Username = flag.String("username", "", "Username used to establish a secure connection with the Jira Rest API")
		conf.Parameters.Password = flag.String("password", "", "Password used by the username in the connection to the Jira Rest API")
		conf.Parameters.Body = flag.String("body", "", "The body of content to set (description, comment, etc.")
		conf.Parameters.Label = flag.String("label", "", "")
		conf.Parameters.IssueID = flag.String("issue-id", "", "")
		conf.Parameters.IssueList = flag.String("issue-list", "", "")
		conf.Parameters.IssueScript = flag.String("issue-script", "", "")

		conf.Initialized = true
	}

	// Parse the flags according to the input parameters
	if !conf.Parsed {
		flag.Parse()
		conf.Parsed = true
	}

	// Setup of the other application flags
	conf.Flags.ApplicationFlags.SingleIssue = *conf.Parameters.IssueID != ""
	conf.Flags.ApplicationFlags.MultipleIssue = *conf.Parameters.IssueList != ""

	// Validations
	successBaseParameters := conf.ValidateBaseParameters()
	successContextParameters := conf.ValidateContextParameters()

	return successBaseParameters && successContextParameters
}
