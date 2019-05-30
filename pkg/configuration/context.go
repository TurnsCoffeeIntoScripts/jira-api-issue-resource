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

	if md.ResourceFlags.SingleIssue {
		c.IssueIds = append(c.IssueIds, *md.ResourceFlags.IssueId)
	} else {
		for _, issue := range strings.Split(*md.ResourceFlags.RawIssueList, ",") {
			c.IssueIds = append(c.IssueIds, issue)
		}
	}

	c.ForceOnParent = *md.ResourceFlags.ForceOnParent
}

func GetExecutionContext(flags JiraApiResourceFlags) *Context {
	ctx := &Context{}
	md := Metadata{}

	md.Initialize(flags)
	ctx.Initialize(md)

	if *flags.CtxComment {
		ctx = SetContextComment(ctx)
	} else if *flags.CtxAddLabel {
		ctx = SetContextAddLabel(ctx)
	}

	return ctx
}
