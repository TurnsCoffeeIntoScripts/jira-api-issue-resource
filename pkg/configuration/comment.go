package configuration

import (
	"net/http"
)

func SetContextComment(ctx *Context) *Context {
	ctx.ApiEndPoint = "issue/" + IssuePlaceholder + "/comment/"
	ctx.HttpMethod = http.MethodPost
	ctx.Body = BuildJsonBodyFromString("body", *ctx.Metadata.ResourceConfiguration.Parameters.Body)

	return ctx
}