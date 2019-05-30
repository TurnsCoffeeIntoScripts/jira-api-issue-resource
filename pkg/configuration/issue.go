package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/domain"
	"net/http"
)

func SetContextAddLabel(ctx *Context) *Context {
	ctx.ApiEndPoint = "issue/" + IssuePlaceholder
	ctx.HttpMethod = http.MethodPut

	issue := domain.AddLabelIssue{}
	issueFields := domain.AddLabelIssueFields{}
	fields := domain.FieldsLabelOnly{}
	labels := make([]domain.JiraLabel, 1)
	label := domain.JiraLabel{Name: *ctx.Metadata.ResourceFlags.Label}

	labels[0] = label
	fields.Labels = labels
	issueFields.Fields = fields
	issue.Update = issueFields

	out, err := json.Marshal(issue)
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return nil
	}

	ctx.Body = out
	return ctx
}

