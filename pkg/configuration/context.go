package configuration

type Context struct {
	IssueId     string
	ApiEndPoint string
	HttpMethod  string
	Metadata    Metadata
	Body        []byte

	// Context parametrization
	ForceOnParent bool
}

func GetExecutionContext(flags JiraApiResourceFlags) Context {
	ctx := Context{}
	md := Metadata{}

	md.Initialize(flags)
	ctx.Metadata = md
	ctx.IssueId = *flags.IssueId
	ctx.ForceOnParent = *flags.ForceOnParent

	if *flags.CtxComment {
		ctx = SetContextComment(ctx)
	}

	return ctx
}
