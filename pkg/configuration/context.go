package configuration

import "strings"

type Context struct {
	IssueIds    []string
	ApiEndPoint string
	HttpMethod  string
	Metadata    Metadata
	Body        []byte

	// Context parametrization
	ForceOnParent bool
}

func (c *Context) Initialize(md Metadata) {
	c.Metadata = md

	if md.ResourceConfiguration.Flags.ApplicationFlags.SingleIssue {
		c.IssueIds = append(c.IssueIds, *md.ResourceConfiguration.Parameters.IssueID)
	} else if md.ResourceConfiguration.Flags.ApplicationFlags.MultipleIssue {
		for _, issue := range strings.Split(*md.ResourceConfiguration.Parameters.IssueList, ",") {
			c.IssueIds = append(c.IssueIds, issue)
		}
	}

	c.ForceOnParent = *md.ResourceConfiguration.Flags.ApplicationFlags.ForceOnParent
}

func GetExecutionContext(conf JiraAPIResourceConfiguration) *Context {
	ctx := &Context{}
	md := Metadata{}

	md.Initialize(conf)
	ctx.Initialize(md)

	if *conf.Flags.ContextFlags.CtxComment.Value {
		ctx = SetContextComment(ctx)
	} else if *conf.Flags.ContextFlags.CtxAddLabel.Value {
		ctx = SetContextAddLabel(ctx)
	}

	return ctx
}
