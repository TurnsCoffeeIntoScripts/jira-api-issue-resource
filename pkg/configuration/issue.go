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

	updatedIssue := domain.UpdatedIssueAddedLabels{}
	addedLabels := domain.AddedLabels{}
	labels := make([]domain.JiraLabel, 1)
	label := domain.JiraLabel{AddedName: *ctx.Metadata.ResourceFlags.Label}

	labels[0] = label
	addedLabels.Labels = labels
	updatedIssue.AddedLabels = addedLabels

	out, err := json.Marshal(updatedIssue)
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return nil
	}

	ctx.Body = out
	return ctx
}

